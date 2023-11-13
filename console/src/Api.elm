module Api exposing
    ( Request(..)
    , Response(..)
    , User
    , request
    )

import Auth
import Http
import Json.Decode as Decode exposing (Decoder)


type Request
    = GetUsersRequest


type alias User =
    { email : String
    , appCodes : List String
    , version : Int
    }


decodeUser : Decoder User
decodeUser =
    Decode.map3 User
        (Decode.field "email" Decode.string)
        (Decode.field "appcodes" <| Decode.list Decode.string)
        (Decode.field "version" Decode.int)


type Response
    = GetUsersResponse (List User)


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
            Op "GET" "/restrictions" Http.emptyBody (Decode.map GetUsersResponse <| Decode.list decodeUser)


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
