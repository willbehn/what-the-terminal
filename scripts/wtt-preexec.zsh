autoload -Uz add-zsh-hook

#NB! Legg til ting du ikke vil skal bli logget her
typeset -a _CLIS_IGNORE
_CLIS_IGNORE+=(
  "* token *" "*apikey*" "*api_key*" "*password*" "*passwd*"
  "*secret*" "*auth*" "*--password*" "*-p*"
  "ssh *" "gpg *" "pass *"
)

_clis_preexec() {
  local line="$1"

  # skipper tomme linjer 
  [[ -z "$line" ]] && return

  # skipper space
  [[ "$line" == \ * ]] && return

  local -a words
  words=(${(z)line})
  local first=$words[1]

  # skipper egne kommandoer
  [[ $first == $WTT_BIN || $first == "wtt" ]] && return

  # skipper sudo kommandoer
  [[ $first == sudo ]] && return

  # ignorer ting som ikke skal bli logget
  local pat
  for pat in "${_CLIS_IGNORE[@]}"; do
    [[ "$line" == $pat ]] && return
  done

  local dir="$PWD"
  local -i ts=$EPOCHSECONDS
  local shell="zsh"

  # json laget med jq pipes til record cmd. Non blocking (&) og gir ingen return verdi (!)
 
  "$WTT_BIN" record \
  --cmd "$line" \
  --dir "$PWD" \
  --ts "$EPOCHSECONDS" \
  --exit "${status:-0}" >/dev/null 2>&1 &!

}

add-zsh-hook preexec _clis_preexec
