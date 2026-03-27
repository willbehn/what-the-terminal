autoload -Uz add-zsh-hook

#NB! Legg til ting du ikke vil skal bli logget her
typeset -a _CLIS_IGNORE
typeset _CLIS_LAST_CMD=""
typeset _CLIS_LAST_DIR=""
typeset -i _CLIS_LAST_TS=0
typeset -i _CLIS_PENDING=0
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
  local first="${words[1]}"

  # ikke logg hvis bin mangler/ikke er executable
  [[ -n "$WTT_BIN" && -x "$WTT_BIN" ]] || return

  # skipper egne kommandoer
  [[ "$first" == "$WTT_BIN" || "$first" == "wtt" ]] && return

  # ignorer ting som ikke skal bli logget
  local pat
  for pat in "${_CLIS_IGNORE[@]}"; do
    [[ "$line" == $pat ]] && return
  done

  _CLIS_LAST_CMD="$line"
  _CLIS_LAST_DIR="$PWD"
  _CLIS_LAST_TS=$EPOCHSECONDS
  _CLIS_PENDING=1
}

_clis_precmd() {
  local -i exit_code=$?

  [[ $_CLIS_PENDING -eq 1 ]] || return
  [[ -n "$WTT_BIN" && -x "$WTT_BIN" ]] || {
    _CLIS_PENDING=0
    return
  }

  # record etter kommandoen er ferdig, slik at exit code blir riktig
  "$WTT_BIN" record \
  --cmd "$_CLIS_LAST_CMD" \
  --dir "$_CLIS_LAST_DIR" \
  --ts "$_CLIS_LAST_TS" \
  --exit "$exit_code" >/dev/null 2>&1 &!

  _CLIS_PENDING=0
  _CLIS_LAST_CMD=""
}

add-zsh-hook preexec _clis_preexec
add-zsh-hook precmd _clis_precmd
