port module Main exposing (main)

import Auth
import Browser
import Browser.Events as Events
import Browser.Navigation as Nav
import Element exposing (Element)
import Element.Lazy as Lazy
import Html exposing (Html)
import Http
import Json.Decode as Decoder exposing (Decoder)
import Route exposing (Route)
import Task
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
    = NOP
    | UrlRequest Browser.UrlRequest
    | UrlChanged Url
    | OnResize Int Int
    | AuthTokenMsg Msg (Result Http.Error Auth.AuthToken)
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
    let
        authModel =
            Auth.init flags.authClientId flags.idp

        authToken =
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
    in
    ( { key = key
      , route = Route.fromUrl url
      , windowSize = ( flags.windowWidth, flags.windowHeight )
      , authToken = authToken
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
      , authModel = authModel
      }
    , Cmd.batch
        [ Nav.pushUrl key (Url.toString url)
        , Maybe.withDefault Cmd.none <|
            Maybe.map
                (\t ->
                    Auth.userInfoRequest UserInfoMsg authModel t.accessToken
                )
                authToken
        ]
    )


maybe : b -> (a -> b) -> Maybe a -> b
maybe default f =
    Maybe.withDefault default << Maybe.map f


maybeCmd : Maybe a -> (a -> Cmd msg) -> Cmd msg
maybeCmd ma f =
    maybe Cmd.none f ma


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
                        , maybeCmd arg.code <|
                            \code ->
                                Auth.tokenRequest (AuthTokenMsg NOP)
                                    model.callbackURL
                                    model.authModel
                                    (Auth.AuthorizationCode code)
                        ]
                    )

                _ ->
                    ( { model | route = route }
                    , Cmd.none
                    )

        OnResize w h ->
            ( { model | windowSize = ( w, h ) }, Cmd.none )

        AuthTokenMsg nextMsg result ->
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
                        [ maybeCmd authToken (storeTokens << t3)
                        , maybeCmd authToken <|
                            \t ->
                                Auth.userInfoRequest UserInfoMsg model.authModel t.accessToken
                        , Task.perform identity <| Task.succeed nextMsg
                        ]
                    )

                Err (Http.BadStatus 401) ->
                    ( model
                    , maybeCmd model.authToken <|
                        \t ->
                            Auth.tokenRequest (AuthTokenMsg nextMsg)
                                model.callbackURL
                                model.authModel
                                (Auth.RefreshToken t.refreshToken)
                    )

                Err _ ->
                    ( { model | authToken = Nothing }, removeTokens () )

        UserInfoMsg res ->
            case res of
                Ok userInfo ->
                    ( { model
                        | authToken =
                            Maybe.map
                                (\t ->
                                    { t | username = Maybe.withDefault preloadUsername userInfo.email }
                                )
                                model.authToken
                      }
                    , Cmd.none
                    )

                Err (Http.BadStatus 401) ->
                    ( model
                    , maybeCmd model.authToken <|
                        \t ->
                            Auth.tokenRequest (AuthTokenMsg msg)
                                model.callbackURL
                                model.authModel
                                (Auth.RefreshToken t.refreshToken)
                    )

                Err _ ->
                    ( model, Cmd.none )

        NOP ->
            ( model, Cmd.none )


subscriptions : Model -> Sub Msg
subscriptions _ =
    Events.onResize OnResize


view : Model -> Browser.Document Msg
view model =
    { title = "Auth console"
    , body =
        [ Element.layout [] <|
            View.Index.view <|
                Lazy.lazy2 View.Org.Header.view
                    model.loginFormURL
                    (Maybe.map (\t -> t.username) model.authToken)
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
