// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package  main

import (
  "github.com/mailru/easyjson/jwriter"
  "github.com/mailru/easyjson/jlexer"
)

func ( Data ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* Data ) UnmarshalJSON([]byte) error { return nil }
func ( Data ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* Data ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_Data *Data

func ( Str ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* Str ) UnmarshalJSON([]byte) error { return nil }
func ( Str ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* Str ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_Str *Str