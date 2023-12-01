module View.UserDeleteConfirm exposing (view)

import Element exposing (Element)
import Element.Font
import Proto.Api
import View.Atom.Button
import View.Atom.DescriptiveList
import View.Template.Form


view :
    { commit : Proto.Api.User -> msg
    , cancel : msg
    }
    -> Proto.Api.User
    -> Element msg
view msgs user =
    View.Template.Form.view
        "Confirm the following user deletions"
        (View.Atom.DescriptiveList.view
            [ { header = "Email"
              , body = Element.text user.email
              }
            , { header = "AppCodes"
              , body = Element.text <| String.join ", " user.appCodes
              }
            ]
        )
        (Element.row [ Element.spacing 20 ]
            [ View.Atom.Button.updateButton (msgs.commit user) "Delete confirm"
            , View.Atom.Button.button msgs.cancel "Cancel"
            ]
        )
