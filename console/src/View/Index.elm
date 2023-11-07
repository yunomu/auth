module View.Index exposing (view)

import Element exposing (Element)
import View.Template.Main


view : Element msg -> Element msg
view header =
    View.Template.Main.view header <|
        Element.text "test"
