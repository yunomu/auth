module View.EditUser exposing
    ( Model
    , Msg
    , init
    , initFromUser
    , update
    , view
    )

import Element exposing (Element)
import Element.Input
import Proto.Api
import Task
import View.Atom.Button
import View.Atom.DescriptiveList
import View.Template.Form


type alias Model =
    { email : String
    , appCodes : List ( Bool, String )
    , version : String
    }


type Msg
    = ChangeEmail String
    | ChangeAppCode Int Bool


init : List String -> Model
init appCodes =
    { email = ""
    , appCodes = List.map (Tuple.pair False) appCodes
    , version = "0"
    }


initFromUser : List String -> Proto.Api.User -> Model
initFromUser appCodes user =
    { email = user.email
    , appCodes = List.map (\appCode -> ( List.member appCode user.appCodes, appCode )) appCodes
    , version = user.version
    }


selectAppCode : (Msg -> msg) -> Int -> ( Bool, String ) -> Element msg
selectAppCode toMsg idx ( checked, appCode ) =
    Element.Input.checkbox []
        { onChange = toMsg << ChangeAppCode idx
        , icon = Element.Input.defaultCheckbox
        , checked = checked
        , label = Element.Input.labelRight [] <| Element.text appCode
        }


listUpdate : Int -> (a -> a) -> List a -> List a
listUpdate idx f ls =
    case ( idx, ls ) of
        ( _, [] ) ->
            []

        ( 0, x :: xs ) ->
            f x :: xs

        ( _, x :: xs ) ->
            x :: listUpdate (idx - 1) f xs


mkUser : Model -> Proto.Api.User
mkUser model =
    { email = model.email
    , appCodes =
        List.filterMap
            (\( checked, appCode ) ->
                if checked then
                    Just appCode

                else
                    Nothing
            )
            model.appCodes
    , version = model.version
    }


view :
    { commit : Proto.Api.User -> msg
    , cancel : msg
    , toMsg : Msg -> msg
    }
    -> Model
    -> Element msg
view msgs model =
    View.Template.Form.view
        "New user"
        (View.Atom.DescriptiveList.view
            [ { header = "Email"
              , body =
                    Element.Input.email []
                        { onChange = msgs.toMsg << ChangeEmail
                        , text = model.email
                        , placeholder = Nothing
                        , label = Element.Input.labelHidden "email"
                        }
              }
            , { header = "AppCodes"
              , body = Element.column [ Element.spacing 7 ] <| List.indexedMap (selectAppCode msgs.toMsg) model.appCodes
              }
            ]
        )
        (Element.row [ Element.spacing 20 ]
            [ View.Atom.Button.updateButton (msgs.commit <| mkUser model) "Commit"
            , View.Atom.Button.button msgs.cancel "Cancel"
            ]
        )


update : Msg -> Model -> ( Model, Cmd msg )
update msg model =
    case msg of
        ChangeEmail s ->
            ( { model | email = s }
            , Cmd.none
            )

        ChangeAppCode i b ->
            ( { model | appCodes = listUpdate i (Tuple.mapFirst (always b)) model.appCodes }
            , Cmd.none
            )
