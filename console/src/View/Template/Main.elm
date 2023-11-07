module View.Template.Main exposing (view)

import Element exposing (Element)


view : Element msg -> Element msg -> Element msg
view header content =
    Element.column
        [ Element.centerX
        , Element.width Element.fill
        , Element.padding 5
        , Element.spacing 20
        ]
        [ header
        , content
        ]
