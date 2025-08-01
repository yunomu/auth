port module Main exposing (main)

import Api
import Auth
import Browser
import Browser.Events as Events
import Browser.Navigation as Nav
import Element exposing (Element)
import Element.Lazy as Lazy
import Html exposing (Html)
import Http
import Json.Decode as Decoder exposing (Decoder)
import Proto.Api
import Route exposing (Route)
import Task
import Url exposing (Url)
import Url.Builder as UrlBuilder
import View.EditProduct
import View.EditUser
import View.Index
import View.Org.Header
import View.ProductDeleteConfirm
import View.Template.Main
import View.UserDeleteConfirm


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
    | UserInfoMsg (Result Http.Error Auth.UserInfo)
    | ApiResponse Api.Request (Result Http.Error Api.Response)
    | AuthResult Msg (Result Http.Error Auth.Token)
    | RedirectToLoginForm
    | RedirectToIndex
    | AuthMsg Msg Auth.Msg
    | EditUserMsg View.EditUser.Msg
    | AddUser
    | AddUserCommit Proto.Api.User
    | EditUser Proto.Api.User
    | EditUserCommit Proto.Api.User
    | DeleteUserConfirm Proto.Api.User
    | DeleteUserCommit Proto.Api.User
    | AddProduct
    | AddProductCommit Proto.Api.Product
    | EditProduct Proto.Api.Product
    | EditProductMsg View.EditProduct.Msg
    | EditProductCommit Proto.Api.Product
    | DeleteProductConfirm Proto.Api.Product
    | DeleteProductCommit Proto.Api.Product
    | RetryUserInfoRequest
    | RetryApiRequest Api.Request


type alias Token =
    { idToken : String
    , accessToken : String
    , refreshToken : String
    }


type alias Model =
    { key : Nav.Key
    , route : Route
    , windowSize : ( Int, Int )
    , username : String
    , loginFormURL : String
    , logoutURL : String
    , logoutRedirectURL : String
    , authModel : Auth.Model
    , endpoint : String
    , headerModel : View.Org.Header.Model
    , users : List Proto.Api.User
    , products : List Proto.Api.Product
    , appCodes : List String
    , editUserModel : View.EditUser.Model
    , editProductModel : View.EditProduct.Model
    }


preloadUsername : String
preloadUsername =
    "me"


init : Flags -> Url -> Nav.Key -> ( Model, Cmd Msg )
init flags url key =
    let
        loginFormURL =
            UrlBuilder.crossOrigin
                flags.idp
                [ "oauth2", "authorize" ]
                [ UrlBuilder.string "response_type" "code"
                , UrlBuilder.string "client_id" flags.authClientId
                , UrlBuilder.string "redirect_uri" flags.authRedirectURL
                ]

        authModel =
            Auth.init flags.authClientId flags.idp flags.authRedirectURL
    in
    ( { key = key
      , route = Route.fromUrl url
      , windowSize = ( flags.windowWidth, flags.windowHeight )
      , username = preloadUsername
      , loginFormURL = loginFormURL
      , logoutURL = UrlBuilder.crossOrigin flags.idp [ "logout" ] []
      , logoutRedirectURL = flags.logoutRedirectURL
      , authModel = authModel
      , endpoint = "/v1"
      , headerModel = View.Org.Header.LoginForm loginFormURL
      , users = []
      , products = []
      , appCodes = []
      , editUserModel = View.EditUser.init []
      , editProductModel = View.EditProduct.init
      }
    , Cmd.batch
        [ Nav.pushUrl key (Url.toString url)
        , Auth.userInfoRequest UserInfoMsg authModel
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
                    , maybeCmd arg.code <|
                        \code ->
                            Auth.tokenRequest RedirectToLoginForm
                                (AuthMsg RedirectToIndex)
                                model.authModel
                                (Auth.AuthorizationCode code)
                    )

                Route.Index ->
                    let
                        getUsers =
                            Api.GetUsersRequest

                        getProducts =
                            Api.GetProductsRequest
                    in
                    ( { model | route = route }
                    , Cmd.batch
                        [ Api.request (ApiResponse getUsers) model.endpoint model.authModel getUsers
                        , Api.request (ApiResponse getProducts) model.endpoint model.authModel getProducts
                        ]
                    )

                _ ->
                    ( { model | route = route }
                    , Cmd.none
                    )

        OnResize w h ->
            ( { model | windowSize = ( w, h ) }, Cmd.none )

        UserInfoMsg res ->
            case res of
                Ok userInfo ->
                    ( { model
                        | headerModel =
                            View.Org.Header.Username <|
                                Maybe.withDefault preloadUsername userInfo.email
                      }
                    , Cmd.none
                    )

                Err (Http.BadStatus 401) ->
                    ( model
                    , Auth.tokenRequest RedirectToLoginForm
                        (AuthMsg RetryUserInfoRequest)
                        model.authModel
                        Auth.RefreshToken
                    )

                Err _ ->
                    ( model, Cmd.none )

        ApiResponse req apiResponse ->
            case apiResponse of
                Ok (Api.GetUsersResponse users) ->
                    ( { model | users = users }
                    , Cmd.none
                    )

                Ok (Api.PostUserResponse user) ->
                    let
                        getUsers =
                            Api.GetUsersRequest
                    in
                    ( { model | route = Route.Index }
                    , Api.request (ApiResponse getUsers) model.endpoint model.authModel getUsers
                    )

                Ok (Api.PutUserResponse user) ->
                    let
                        getUsers =
                            Api.GetUsersRequest
                    in
                    ( { model | route = Route.Index }
                    , Api.request (ApiResponse getUsers) model.endpoint model.authModel getUsers
                    )

                Ok (Api.DeleteUserResponse email) ->
                    let
                        getUsers =
                            Api.GetUsersRequest
                    in
                    ( { model | route = Route.Index }
                    , Api.request (ApiResponse getUsers) model.endpoint model.authModel getUsers
                    )

                Ok (Api.GetProductsResponse products) ->
                    ( { model
                        | products = products
                        , appCodes = List.map (\p -> p.appCode) products
                      }
                    , Cmd.none
                    )

                Ok (Api.PostProductResponse product) ->
                    let
                        getProducts =
                            Api.GetProductsRequest
                    in
                    ( { model | route = Route.Index }
                    , Api.request (ApiResponse getProducts) model.endpoint model.authModel getProducts
                    )

                Ok (Api.PutProductResponse product) ->
                    let
                        getProducts =
                            Api.GetProductsRequest
                    in
                    ( { model | route = Route.Index }
                    , Api.request (ApiResponse getProducts) model.endpoint model.authModel getProducts
                    )

                Ok (Api.DeleteProductResponse clientId) ->
                    let
                        getProducts =
                            Api.GetProductsRequest
                    in
                    ( { model | route = Route.Index }
                    , Api.request (ApiResponse getProducts) model.endpoint model.authModel getProducts
                    )

                Err (Http.BadStatus 401) ->
                    ( model
                    , Auth.tokenRequest RedirectToLoginForm
                        (AuthMsg (RetryApiRequest req))
                        model.authModel
                        Auth.RefreshToken
                    )

                Err err ->
                    ( model, Cmd.none )

        AuthResult prevMsg result ->
            case result of
                Ok token ->
                    ( model
                    , Cmd.batch
                        [ storeTokens ( token.idToken, token.accessToken, token.refreshToken )
                        , Task.perform (\_ -> prevMsg) (Task.succeed ())
                        ]
                    )

                Err err ->
                    ( model, Cmd.none )

        RedirectToLoginForm ->
            ( model, Nav.load model.loginFormURL )

        AuthMsg prevMsg authMsg ->
            let
                ( authModel, cmd ) =
                    Auth.update (AuthResult prevMsg) authMsg model.authModel
            in
            ( { model | authModel = authModel }
            , cmd
            )

        RedirectToIndex ->
            ( model
            , Nav.pushUrl model.key <| Route.path Route.Index
            )

        AddProduct ->
            ( { model
                | route = Route.ProductAdd
                , editProductModel = View.EditProduct.init
              }
            , Cmd.none
            )

        AddUser ->
            ( { model
                | route = Route.UserAdd
                , editUserModel = View.EditUser.init model.appCodes
              }
            , Cmd.none
            )

        AddUserCommit user ->
            let
                req =
                    Api.PostUserRequest user
            in
            ( model
            , Api.request (ApiResponse req) model.endpoint model.authModel req
            )

        EditUser user ->
            ( { model
                | route = Route.UserEdit
                , editUserModel = View.EditUser.initFromUser model.appCodes user
              }
            , Cmd.none
            )

        EditUserCommit user ->
            let
                req =
                    Api.PutUserRequest user
            in
            ( model
            , Api.request (ApiResponse req) model.endpoint model.authModel req
            )

        DeleteUserConfirm user ->
            ( { model | route = Route.UserDeleteConfirm user }
            , Cmd.none
            )

        DeleteUserCommit user ->
            let
                req =
                    Api.DeleteUserRequest user.email
            in
            ( model
            , Api.request (ApiResponse req) model.endpoint model.authModel req
            )

        EditUserMsg msg_ ->
            let
                ( editUserModel, cmd ) =
                    View.EditUser.update msg_ model.editUserModel
            in
            ( { model | editUserModel = editUserModel }
            , cmd
            )

        DeleteProductConfirm product ->
            ( { model | route = Route.ProductDeleteConfirm product }
            , Cmd.none
            )

        DeleteProductCommit product ->
            let
                req =
                    Api.DeleteProductRequest product.clientId
            in
            ( model
            , Api.request (ApiResponse req) model.endpoint model.authModel req
            )

        EditProductMsg msg_ ->
            let
                ( editProductModel, cmd ) =
                    View.EditProduct.update msg_ model.editProductModel
            in
            ( { model | editProductModel = editProductModel }
            , cmd
            )

        AddProductCommit product ->
            let
                req =
                    Api.PostProductRequest product
            in
            ( model
            , Api.request (ApiResponse req) model.endpoint model.authModel req
            )

        EditProduct product ->
            ( { model
                | route = Route.ProductEdit
                , editProductModel = View.EditProduct.initFromProduct product
              }
            , Cmd.none
            )

        EditProductCommit product ->
            let
                req =
                    Api.PutProductRequest product
            in
            ( model
            , Api.request (ApiResponse req) model.endpoint model.authModel req
            )

        RetryUserInfoRequest ->
            ( model, Auth.userInfoRequest UserInfoMsg model.authModel )

        RetryApiRequest req ->
            ( model, Api.request (ApiResponse req) model.endpoint model.authModel req )

        NOP ->
            ( model, Cmd.none )


subscriptions : Model -> Sub Msg
subscriptions _ =
    Events.onResize OnResize


indexView =
    View.Index.view
        { addUser = AddUser
        , deleteUser = DeleteUserConfirm
        , editUser = EditUser
        , addProduct = AddProduct
        , deleteProduct = DeleteProductConfirm
        , editProduct = EditProduct
        }


userDeleteView =
    View.UserDeleteConfirm.view
        { commit = DeleteUserCommit
        , cancel = RedirectToIndex
        }


addUserView =
    View.EditUser.view
        { commit = AddUserCommit
        , cancel = RedirectToIndex
        , toMsg = EditUserMsg
        }


editUserView =
    View.EditUser.view
        { commit = EditUserCommit
        , cancel = RedirectToIndex
        , toMsg = EditUserMsg
        }


addProductView =
    View.EditProduct.view "New product"
        { commit = AddProductCommit
        , cancel = RedirectToIndex
        , toMsg = EditProductMsg
        }


productDeleteView =
    View.ProductDeleteConfirm.view
        { commit = DeleteProductCommit
        , cancel = RedirectToIndex
        }


editProductView =
    View.EditProduct.view "Edit product"
        { commit = EditProductCommit
        , cancel = RedirectToIndex
        , toMsg = EditProductMsg
        }


view : Model -> Browser.Document Msg
view model =
    { title = "Auth console"
    , body =
        let
            header =
                Lazy.lazy View.Org.Header.view model.headerModel
        in
        [ Element.layout [] <|
            View.Template.Main.view header <|
                case model.route of
                    Route.Index ->
                        Lazy.lazy2 indexView model.users model.products

                    Route.UserAdd ->
                        Lazy.lazy addUserView model.editUserModel

                    Route.UserEdit ->
                        Lazy.lazy editUserView model.editUserModel

                    Route.UserDeleteConfirm user ->
                        Lazy.lazy userDeleteView user

                    Route.ProductAdd ->
                        Lazy.lazy addProductView model.editProductModel

                    Route.ProductEdit ->
                        Lazy.lazy editProductView model.editProductModel

                    Route.ProductDeleteConfirm product ->
                        Lazy.lazy productDeleteView product

                    Route.AuthCallback _ ->
                        Element.none

                    Route.NotFound _ ->
                        Element.text "NotFound"
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
