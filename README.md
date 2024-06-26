#Emulate generation of VTT subitles.

Usage of ./test_subs: 
  -delay int
        Specify a delay in seconds between each subtitle, (default 10)
  -fifoname string
        Specify a name for mkfifo named pipe, 
                        default is 'subtitle_pipe' in the current folder (default "/tmp/subtitle_pipe")
  -inc
        If true, timecodes in subtitles will incrementing to textduration, default is true.
  -mode string
        Specify a mode of writing subtitles:
                         'file' writes each sub to a separate file, 
                         'stdout' writes each to stdout, 
                         'udp' writes each to udp connection, see udpport setting
                         'pipe' writes to mkfifo, see fifoname setting. 
                         'listen'  runs UDP listener to debug UDP output
                         Default is 'file' (default "file")
  -offset int
        Specify an offset for timecodes, default is 0
  -text string
        User defined text for subtitles, default is number of current subtitle.
                        You can use {s} to substitute current time and {n} for current subtitle number
  -textdur int
        Specify a text duration in subtitle, default is 10.
  -udpport int
        Specify a port for udp (default 9723)

If you run in mode "listen" , then the program doesn't producing any output. It is listening specified UDP port (default 9273).
You should run in separate terminal window in mode "udp" to produce some content.
