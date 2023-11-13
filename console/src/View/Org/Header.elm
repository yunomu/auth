module View.Org.Header exposing
    ( Model(..)
    , view
    )

import Element exposing (Element)
import Element.Border as Border


type Model
    = LoginForm String
    | Username String


edges =
    { top = 0
    , bottom = 0
    , left = 0
    , right = 0
    }


viewUser : Model -> Element msg
viewUser model =
    Element.el [ Element.alignRight ] <|
        case model of
            Username u ->
                Element.text <| "User: " ++ u

            LoginForm url ->
                Element.link []
                    { url = url
                    , label = Element.text "Login/Signup"
                    }


view : Model -> Element msg
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
        , viewUser model
        ]
