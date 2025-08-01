module Api exposing
    ( Request(..)
    , Response(..)
    , request
    )

import Auth
import Http
import Json.Decode as Decode exposing (Decoder)
import Json.Encode as Encode
import Proto.Api
import Url.Builder as B


type Request
    = GetUsersRequest
    | PostUserRequest Proto.Api.User
    | PutUserRequest Proto.Api.User
    | DeleteUserRequest String
    | GetProductsRequest
    | PostProductRequest Proto.Api.Product
    | PutProductRequest Proto.Api.Product
    | DeleteProductRequest String


type Response
    = GetUsersResponse (List Proto.Api.User)
    | PostUserResponse Proto.Api.User
    | PutUserResponse Proto.Api.User
    | DeleteUserResponse String
    | GetProductsResponse (List Proto.Api.Product)
    | PostProductResponse Proto.Api.Product
    | PutProductResponse Proto.Api.Product
    | DeleteProductResponse String


type ResponseDecoder
    = Json (Decoder Response)
    | Empty Response


type alias Op =
    { method : String
    , path : String
    , body : Http.Body
    , decoder : ResponseDecoder
    }


mkOp : Request -> Op
mkOp req =
    case req of
        GetUsersRequest ->
            { method = "GET"
            , path = B.absolute [ "restrictions" ] []
            , body = Http.emptyBody
            , decoder =
                Json <| Decode.map (\r -> GetUsersResponse r.users) Proto.Api.restrictionsResponseDecoder
            }

        PostUserRequest user ->
            { method = "POST"
            , path = B.absolute [ "restrictions" ] []
            , body = Http.jsonBody <| Proto.Api.userEncoder user
            , decoder =
                Json <|
                    Decode.andThen
                        (\r -> Decode.succeed (PostUserResponse r.user))
                        Proto.Api.restrictionResponseDecoder
            }

        PutUserRequest user ->
            { method = "PUT"
            , path = B.absolute [ "restrictions", user.email ] []
            , body = Http.jsonBody <| Proto.Api.userEncoder user
            , decoder = Empty (PutUserResponse user)
            }

        DeleteUserRequest id ->
            { method = "DELETE"
            , path = B.absolute [ "restrictions", id ] []
            , body = Http.emptyBody
            , decoder = Empty (DeleteUserResponse id)
            }

        GetProductsRequest ->
            { method = "GET"
            , path = B.absolute [ "products" ] []
            , body = Http.emptyBody
            , decoder =
                Json <|
                    Decode.map (\r -> GetProductsResponse r.products) <|
                        Proto.Api.productsResponseDecoder
            }

        PostProductRequest product ->
            { method = "POST"
            , path = B.absolute [ "products" ] []
            , body = Http.jsonBody <| Proto.Api.productEncoder product
            , decoder =
                Json <|
                    Decode.andThen
                        (\r -> Decode.succeed (PostProductResponse r.product))
                        Proto.Api.productResponseDecoder
            }

        PutProductRequest product ->
            { method = "PUT"
            , path = B.absolute [ "products", product.clientId ] []
            , body = Http.jsonBody <| Proto.Api.productEncoder product
            , decoder = Empty (PutProductResponse product)
            }

        DeleteProductRequest id ->
            { method = "DELETE"
            , path = B.absolute [ "products", id ] []
            , body = Http.emptyBody
            , decoder = Empty (DeleteProductResponse id)
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
        , expect =
            case op.decoder of
                Json decoder ->
                    Http.expectJson toMsg decoder

                Empty res ->
                    Http.expectString (toMsg << Result.map (always res))
        }
