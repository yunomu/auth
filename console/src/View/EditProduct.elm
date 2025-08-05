module View.EditProduct exposing (Model, Msg, init, initFromProduct, update, view)

import Element exposing (Element)
import Element.Input
import Proto.Api
import Task
import View.Atom.Button
import View.Atom.DescriptiveList
import View.Template.Form


type alias Model =
    { clientId : String
    , appCode : String
    , funcArn : String
    , version : String
    }


type Msg
    = ChangeClientId String
    | ChangeAppCode String
    | ChangeFuncArn String


init : Model
init =
    { clientId = ""
    , appCode = ""
    , funcArn = ""
    , version = "0"
    }


initFromProduct : Proto.Api.Product -> Model
initFromProduct product =
    { clientId = product.clientId
    , appCode = product.appCode
    , funcArn = Maybe.withDefault "" product.funcArn
    , version = product.version
    }


mk : Model -> Proto.Api.Product
mk model =
    { clientId = model.clientId
    , appCode = model.appCode
    , funcArn =
        if model.funcArn == "" then
            Nothing

        else
            Just model.funcArn
    , version = model.version
    }


view :
    String
    ->
        { commit : Proto.Api.Product -> msg
        , cancel : msg
        , toMsg : Msg -> msg
        }
    -> Model
    -> Element msg
view subject msgs model =
    View.Template.Form.view
        subject
        (View.Atom.DescriptiveList.view
            [ { header = "ClientId"
              , body =
                    Element.Input.email []
                        { onChange = msgs.toMsg << ChangeClientId
                        , text = model.clientId
                        , placeholder = Nothing
                        , label = Element.Input.labelHidden "clientId"
                        }
              }
            , { header = "AppCode"
              , body =
                    Element.Input.email []
                        { onChange = msgs.toMsg << ChangeAppCode
                        , text = model.appCode
                        , placeholder = Nothing
                        , label = Element.Input.labelHidden "appCode"
                        }
              }
            , { header = "FuncArn"
              , body =
                    Element.Input.email []
                        { onChange = msgs.toMsg << ChangeFuncArn
                        , text = model.funcArn
                        , placeholder = Nothing
                        , label = Element.Input.labelHidden "funcArn"
                        }
              }
            ]
        )
        (Element.row [ Element.spacing 20 ]
            [ View.Atom.Button.updateButton (msgs.commit <| mk model) "Commit"
            , View.Atom.Button.button msgs.cancel "Cancel"
            ]
        )


update : Msg -> Model -> ( Model, Cmd msg )
update msg model =
    case msg of
        ChangeClientId s ->
            ( { model | clientId = s }
            , Cmd.none
            )

        ChangeAppCode s ->
            ( { model | appCode = s }
            , Cmd.none
            )

        ChangeFuncArn s ->
            ( { model | funcArn = s }
            , Cmd.none
            )
