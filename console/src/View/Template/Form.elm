module View.Template.Form exposing (view)

import Element exposing (Element)


view : String -> Element msg -> Element msg -> Element msg
view caption form operation =
    Element.column [ Element.paddingXY 20 0, Element.spacing 30 ]
        [ Element.text caption
        , form
        , operation
        ]
