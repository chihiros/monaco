{{ block "mona_front" . }}
 ∧___∧   /{{ printOverLine . }}
( ´∀ `) < {{ echo . -}}
(     )  \{{ printUnderLine . }}
|  |  |
(__)__)
{{ end }}

{{ block "mona_back" . }}
 ∧___∧   /{{ printOverLine . }}
(     ) < {{ . }}
(  O  )  \{{ printUnderLine . }}
|  |  |
(__)__)
{{ end }}
