module View.Org.Header exposing (view)

import Element exposing (Element)
import Element.Border as Border


edges =
    { top = 0
    , bottom = 0
    , left = 0
    , right = 0
    }


viewUser : String -> Maybe String -> Element msg
viewUser loginUrl username =
    Element.el [ Element.alignRight ] <|
        case username of
            Just u ->
                Element.text <| "User: " ++ u

            Nothing ->
                Element.link []
                    { url = loginUrl
                    , label = Element.text "Login/Signup"
                    }


view :
    String
    -> Maybe String
    -> Element msg
view loginUrl username =
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
        , viewUser loginUrl username
        ]
