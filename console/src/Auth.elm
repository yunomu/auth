module Auth exposing
    ( AuthToken
    , Model
    , TokenRequestType(..)
    , UserInfo
    , init
    , tokenRequest
    , userInfoRequest
    )

import Http
import Json.Decode as Decoder exposing (Decoder)
import Url.Builder as UrlBuilder


type alias Model =
    { clientId : String
    , tokenEndpoint : String
    , userInfoEndpoint : String
    }


init : String -> String -> Model
init clientId idp =
    { clientId = clientId
    , tokenEndpoint = UrlBuilder.crossOrigin idp [ "oauth2", "token" ] []
    , userInfoEndpoint = UrlBuilder.crossOrigin idp [ "oauth2", "userInfo" ] []
    }


type alias AuthToken =
    { idToken : String
    , accessToken : String
    , refreshToken : Maybe String
    , expiresIn : Int
    , tokenType : String
    }


authTokenDecoder : Decoder AuthToken
authTokenDecoder =
    Decoder.map5 AuthToken
        (Decoder.field "id_token" Decoder.string)
        (Decoder.field "access_token" Decoder.string)
        (Decoder.maybe <| Decoder.field "refresh_token" Decoder.string)
        (Decoder.field "expires_in" Decoder.int)
        (Decoder.field "token_type" Decoder.string)


type TokenRequestType
    = AuthorizationCode String
    | RefreshToken String


tokenRequest :
    (Result Http.Error AuthToken -> msg)
    -> String
    -> Model
    -> TokenRequestType
    -> Cmd msg
tokenRequest authTokenMsg redirectUri model grantType =
    let
        addGrantType ls =
            case grantType of
                AuthorizationCode code ->
                    ( "grant_type", "authorization_code" )
                        :: ( "code", code )
                        :: ls

                RefreshToken token ->
                    ( "grant_type", "refresh_token" )
                        :: ( "refresh_token", token )
                        :: ls
    in
    Http.post
        { url = model.tokenEndpoint
        , body =
            Http.stringBody "application/x-www-form-urlencoded" <|
                String.join "&" <|
                    List.map (\( a, b ) -> String.concat [ a, "=", b ]) <|
                        addGrantType
                            [ ( "client_id", model.clientId )
                            , ( "redirect_uri", redirectUri )
                            ]
        , expect = Http.expectJson authTokenMsg authTokenDecoder
        }


type alias UserInfo =
    { sub : String
    , name : Maybe String
    , email : Maybe String
    }


userInfoDecoder : Decoder UserInfo
userInfoDecoder =
    Decoder.map3 UserInfo
        (Decoder.field "sub" Decoder.string)
        (Decoder.maybe <| Decoder.field "name" Decoder.string)
        (Decoder.maybe <| Decoder.field "email" Decoder.string)


userInfoRequest :
    (Result Http.Error UserInfo -> msg)
    -> Model
    -> String
    -> Cmd msg
userInfoRequest userInfoMsg model accessToken =
    Http.request
        { method = "POST"
        , headers = [ Http.header "Authorization" <| "Bearer " ++ accessToken ]
        , url = model.userInfoEndpoint
        , body = Http.emptyBody
        , expect = Http.expectJson userInfoMsg userInfoDecoder
        , timeout = Nothing
        , tracker = Nothing
        }
