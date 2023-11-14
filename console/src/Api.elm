module Api exposing
    ( Request(..)
    , Response(..)
    , request
    )

import Auth
import Http
import Json.Decode as Decode exposing (Decoder)
import Proto.Api


type Request
    = GetUsersRequest


type Response
    = GetUsersResponse (List Proto.Api.User)


type alias Op =
    { method : String
    , path : String
    , body : Http.Body
    , decoder : Decoder Response
    }


mkOp : Request -> Op
mkOp req =
    case req of
        GetUsersRequest ->
            { method = "GET"
            , path = "/restrictions"
            , body = Http.emptyBody
            , decoder =
                Decode.map (\r -> GetUsersResponse r.users) <|
                    Proto.Api.restrictionsResponseDecoder
            }


maybe : b -> (a -> b) -> Maybe a -> b
maybe b f =
    Maybe.withDefault b << Maybe.map f


request : (Result Http.Error Response -> msg) -> String -> Auth.Model -> Request -> Cmd msg
request toMsg endpoint authModel req =
    let
        op =
            mkOp req
    in
    Auth.signedRequest authModel
        { method = op.method
        , headers = []
        , url = endpoint ++ op.path
        , body = op.body
        , expect = Http.expectJson toMsg op.decoder
        }
