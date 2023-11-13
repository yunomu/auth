module View.Index exposing (view)

import Api
import Element exposing (Element)
import View.Atom.Table
import View.Template.Main


view : Element msg -> List Api.User -> Element msg
view header users =
    View.Template.Main.view header <|
        if List.isEmpty users then
            Element.none

        else
            View.Atom.Table.view users
                (\_ _ -> False)
                [ { header = "Email"
                  , view = \i user -> Element.text user.email
                  }
                , { header = "AppCodes"
                  , view = \i user -> Element.text <| String.join ", " user.appCodes
                  }
                ]
