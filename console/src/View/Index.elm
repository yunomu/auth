module View.Index exposing (view)

import Api
import Element exposing (Attribute, Element)
import Element.Background
import Element.Border
import Element.Font
import Element.Input
import Element.Lazy as Lazy
import Proto.Api
import View.Atom.Button
import View.Atom.Table


section : String -> Element msg -> Element msg
section label elem =
    Element.column [ Element.spacing 10 ]
        [ Element.el [] <| Element.text label
        , elem
        ]


crudTable :
    msg
    -> (record -> msg)
    -> (record -> msg)
    -> String
    -> List record
    ->
        List
            { header : String
            , view : Int -> record -> Element msg
            }
    -> Element msg
crudTable addMsg delMsg editMsg label data columns =
    Element.column [ Element.spacing 10 ]
        [ Element.el [] <| Element.text label
        , View.Atom.Table.view data
            (\_ _ -> False)
            (columns
                ++ [ { header = ""
                     , view =
                        \i r ->
                            Element.row [ Element.spacing 5 ]
                                [ View.Atom.Button.button (editMsg r) "Edit"
                                , View.Atom.Button.updateButton (delMsg r) "Delete"
                                ]
                     }
                   ]
            )
        , View.Atom.Button.button addMsg "Add"
        ]


usersTable : msg -> (Proto.Api.User -> msg) -> (Proto.Api.User -> msg) -> List Proto.Api.User -> Element msg
usersTable addUser deleteUser editUser users =
    crudTable addUser deleteUser editUser "Users" users <|
        [ { header = "Email"
          , view = \i user -> Element.text user.email
          }
        , { header = "AppCodes"
          , view = \i user -> Element.text <| String.join ", " user.appCodes
          }
        ]


productsTable : msg -> (Proto.Api.Product -> msg) -> (Proto.Api.Product -> msg) -> List Proto.Api.Product -> Element msg
productsTable addProduct deleteProduct editProduct products =
    crudTable addProduct deleteProduct editProduct "Products" products <|
        [ { header = "Client ID"
          , view = \i product -> Element.text product.clientId
          }
        , { header = "AppCode"
          , view = \i product -> Element.text product.appCode
          }
        , { header = "Func ARN"
          , view = \i product -> Element.text <| Maybe.withDefault "---" product.funcArn
          }
        ]


view :
    { addUser : msg
    , deleteUser : Proto.Api.User -> msg
    , editUser : Proto.Api.User -> msg
    , addProduct : msg
    , deleteProduct : Proto.Api.Product -> msg
    , editProduct : Proto.Api.Product -> msg
    }
    -> List Proto.Api.User
    -> List Proto.Api.Product
    -> Element msg
view msgs users products =
    Element.column [ Element.paddingXY 10 0, Element.spacing 20 ]
        [ Lazy.lazy4 usersTable msgs.addUser msgs.deleteUser msgs.editUser users
        , Lazy.lazy4 productsTable msgs.addProduct msgs.deleteProduct msgs.editProduct products
        ]
