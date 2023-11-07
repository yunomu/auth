module Route exposing (Route(..), fromUrl, path)

import Url exposing (Url)
import Url.Parser as P exposing ((<?>), Parser, s)
import Url.Parser.Query as Query


type Route
    = Index
    | AuthCallback
        { code : Maybe String
        , state : Maybe String
        }
    | NotFound Url


authCallback : Maybe String -> Maybe String -> Route
authCallback code state =
    AuthCallback { code = code, state = state }


parser : Parser (Route -> a) a
parser =
    P.oneOf
        [ P.map Index P.top
        , P.map authCallback <| s "callback" <?> Query.string "code" <?> Query.string "state"
        ]


path : Route -> String
path route =
    "/"


fromUrl : Url -> Route
fromUrl url =
    Maybe.withDefault (NotFound url) <| P.parse parser url
