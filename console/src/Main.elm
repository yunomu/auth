port module Main exposing (main)

import Auth
import Browser
import Browser.Events as Events
import Browser.Navigation as Nav
import Element exposing (Element)
import Html exposing (Html)
import Http
import Json.Decode as Decoder exposing (Decoder)
import Route exposing (Route)
import Url exposing (Url)
import Url.Builder as UrlBuilder
import View.Index
import View.Org.Header


port storeTokens : ( String, String, String ) -> Cmd msg


port removeTokens : () -> Cmd msg


type alias Flags =
    { idToken : Maybe String
    , accessToken : Maybe String
    , refreshToken : Maybe String
    , windowWidth : Int
    , windowHeight : Int
    , authClientId : String
    , authRedirectURL : String
    , logoutRedirectURL : String
    , idp : String
    }


type Msg
    = UrlRequest Browser.UrlRequest
    | UrlChanged Url
    | OnResize Int Int
    | AuthTokenMsg (Result Http.Error Auth.AuthToken)
    | UserInfoMsg (Result Http.Error Auth.UserInfo)


type alias Token =
    { idToken : String
    , accessToken : String
    , refreshToken : String
    , username : String
    }


type alias Model =
    { key : Nav.Key
    , route : Route
    , windowSize : ( Int, Int )
    , authToken : Maybe Token
    , callbackURL : String
    , loginFormURL : String
    , logoutURL : String
    , logoutRedirectURL : String
    , authModel : Auth.Model
    }


preloadUsername : String
preloadUsername =
    "me"


init : Flags -> Url -> Nav.Key -> ( Model, Cmd Msg )
init flags url key =
    ( { key = key
      , route = Route.fromUrl url
      , windowSize = ( flags.windowWidth, flags.windowHeight )
      , authToken =
            case ( flags.idToken, flags.accessToken, flags.refreshToken ) of
                ( Just idToken, Just accessToken, Just refreshToken ) ->
                    Just
                        { idToken = idToken
                        , accessToken = accessToken
                        , refreshToken = refreshToken
                        , username = preloadUsername
                        }

                _ ->
                    Nothing
      , callbackURL = flags.authRedirectURL
      , loginFormURL =
            UrlBuilder.crossOrigin
                flags.idp
                [ "oauth2", "authorize" ]
                [ UrlBuilder.string "response_type" "code"
                , UrlBuilder.string "client_id" flags.authClientId
                , UrlBuilder.string "redirect_uri" flags.authRedirectURL
                ]
      , logoutURL = UrlBuilder.crossOrigin flags.idp [ "logout" ] []
      , logoutRedirectURL = flags.logoutRedirectURL
      , authModel = Auth.init flags.authClientId flags.idp
      }
    , Nav.pushUrl key (Url.toString url)
    )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        UrlRequest urlRequest ->
            case urlRequest of
                Browser.Internal url ->
                    ( model, Nav.pushUrl model.key (Url.toString url) )

                Browser.External href ->
                    ( model, Nav.load href )

        UrlChanged url ->
            let
                route =
                    Route.fromUrl url
            in
            case route of
                Route.AuthCallback arg ->
                    ( model
                    , Cmd.batch
                        [ Nav.pushUrl model.key <| Route.path Route.Index
                        , case arg.code of
                            Just code ->
                                Auth.tokenRequest AuthTokenMsg
                                    model.callbackURL
                                    model.authModel
                                    (Auth.AuthorizationCode code)

                            Nothing ->
                                Cmd.none
                        ]
                    )

                _ ->
                    ( { model | route = route }
                    , Cmd.none
                    )

        OnResize w h ->
            ( { model | windowSize = ( w, h ) }, Cmd.none )

        AuthTokenMsg result ->
            case result of
                Ok res ->
                    let
                        authToken =
                            case ( res.refreshToken, model.authToken ) of
                                ( Just refreshToken, mt ) ->
                                    Just
                                        { idToken = res.idToken
                                        , accessToken = res.accessToken
                                        , refreshToken = refreshToken
                                        , username =
                                            Maybe.withDefault preloadUsername <|
                                                Maybe.map (\t -> t.username) mt
                                        }

                                ( _, Just authToken_ ) ->
                                    Just
                                        { authToken_
                                            | idToken = res.idToken
                                            , accessToken = res.accessToken
                                        }

                                _ ->
                                    -- not reached
                                    -- Refresh token requested when there is no refresh token
                                    Nothing

                        t3 t =
                            ( t.idToken, t.accessToken, t.refreshToken )
                    in
                    ( { model | authToken = authToken }
                    , Cmd.batch
                        [ Maybe.withDefault Cmd.none <| Maybe.map (storeTokens << t3) authToken
                        , Maybe.withDefault Cmd.none <|
                            Maybe.map
                                (\t ->
                                    Auth.userInfoRequest UserInfoMsg model.authModel t.accessToken
                                )
                                authToken
                        ]
                    )

                Err _ ->
                    ( model, Cmd.none )

        UserInfoMsg res ->
            case res of
                Ok userInfo ->
                    let
                        _ =
                            Debug.log "userInfo" userInfo

                        authToken =
                            Maybe.map
                                (\t ->
                                    { t | username = Maybe.withDefault preloadUsername userInfo.email }
                                )
                                model.authToken
                    in
                    ( { model | authToken = authToken }
                    , Cmd.none
                    )

                Err _ ->
                    ( model, Cmd.none )


subscriptions : Model -> Sub Msg
subscriptions _ =
    Events.onResize OnResize


view : Model -> Browser.Document Msg
view model =
    { title = "Login test"
    , body =
        [ Element.layout [] <|
            View.Index.view <|
                View.Org.Header.view
                    { loginUrl = model.loginFormURL
                    }
        ]
    }


main : Program Flags Model Msg
main =
    Browser.application
        { init = init
        , update = update
        , subscriptions = subscriptions
        , view = view
        , onUrlChange = UrlChanged
        , onUrlRequest = UrlRequest
        }
