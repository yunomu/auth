module View.ProductDeleteConfirm exposing (view)

import Element exposing (Element)
import Element.Font
import Proto.Api
import View.Atom.Button
import View.Atom.DescriptiveList
import View.Template.Form


view :
    { commit : Proto.Api.Product -> msg
    , cancel : msg
    }
    -> Proto.Api.Product
    -> Element msg
view msgs product =
    View.Template.Form.view
        "Confirm the following product deletions"
        (View.Atom.DescriptiveList.view
            [ { header = "Client ID"
              , body = Element.text product.clientId
              }
            , { header = "AppCode"
              , body = Element.text product.appCode
              }
            , { header = "Function ARN"
              , body = Element.text product.funcArn
              }
            ]
        )
        (Element.row [ Element.spacing 20 ]
            [ View.Atom.Button.updateButton (msgs.commit product) "Delete confirm"
            , View.Atom.Button.button msgs.cancel "Cancel"
            ]
        )
