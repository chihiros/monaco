{{ block "mona_front" . }}
 ∧___∧   /{{ printOverLine . }}
( ´∀ `) < {{ echo . -}}
(     )  \{{ printUnderLine . }}
|  |  |
(__)__)
{{ end }}

{{ block "mona_back" . }}
 ∧___∧   /{{ printOverLine . }}
(     ) < {{ echo . -}}
(  O  )  \{{ printUnderLine . }}
|  |  |
(__)__)
{{ end }}
