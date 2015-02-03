# GOlang-Tool-fix-subtitle
It's a great tool for fix delay/hurry subtitle. Program can shift time to forward or back, for all expressions.

Run tool

    go run subtitle.go

Where [flags] are:<br>
  -f="": [required] Path to file with subtitles<br>
  -n="": New output file. Default name will be the same file name "-f" with prefix _<br>
  -o="": Set "true" for output on display<br>
  -s="": [required] Seconds, shift all timeline. It's should be integer, or negative integer<br>

Move subtitles forward on +5 seconds

    go run subtitle.go -f "move.srt" -s 5

Move subtitles back on -5 seconds

    go run subtitle.go -f "move.srt" -s -5

Create new file with subtitles 

    go run subtitle.go -f "move.srt" -s 5 -n "fixed.srt"

Show output 

    go run subtitle.go -f "move.srt" -s 5 -o true
