module View.Org.Header exposing (view)

import Element exposing (Element)
import Element.Border as Border


edges =
    { top = 0
    , bottom = 0
    , left = 0
    , right = 0
    }


view : { loginUrl : String } -> Element msg
view model =
    Element.row
        [ Element.width Element.fill
        , Element.padding 5
        , Border.widthEach { edges | bottom = 1 }
        ]
        [ Element.el
            [ Element.alignLeft
            ]
          <|
            Element.text "Auth console"
        , Element.link [ Element.alignRight ]
            { url = model.loginUrl
            , label = Element.text "Login/Signup"
            }
        ]
