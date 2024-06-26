# Emulate generation of VTT subitles.

## Command line parameters

*-delay*
        Int, Specify a delay in seconds between each subtitle, (default 10)
        
*-fifoname* 
        String, Specify a name for mkfifo named pipe, 
        default is 'subtitle_pipe' in the current folder (default "/tmp/subtitle_pipe")
                        
*-inc*
        Boolean, If true, timecodes in subtitles will incrementing to textduration, default is true.


*-mode*
     String, specify a mode of writing subtitles:
* 'file' writes each sub to a separate file, 
* 'stdout' writes each to stdout, 
* 'udp' writes each to udp connection, see udpport setting
* 'pipe' writes to mkfifo, see fifoname setting. 
* 'listen'  runs UDP listener to debug UDP output                        
Default is "file"

                         
*-offset*
     Int, Specify an offset for timecodes, in seconds, default is 0

        
*-text*
        String, User defined text for subtitles, default is number of current subtitle.
        You can use {s} to substitute current time and {n} for current subtitle number

                        
*-textdur* 
       Int, Specify a text duration in subtitle in seconds, default is 10.

        
*-udpport*
        Int, Specify a port for udp (default 9723)
       

If you run in mode "listen" , then the program doesn't produce any output. Instead, it is listening specified UDP port (default 9273).
You should run another instance of the program in the separate terminal window in mode "udp" to produce some content.

## Examples

Generate consequent subtitles for timecodes 0-10 sec, each 10 seconds, timecodes are incrementing, writing to pipe:

`./test_subs  -mode=pipe `

Generate sequence of subtitles using udp

`./test_subs  -mode=udp`

Generate using udp each 3 seconds, timecodes will be: 0-3, 3-6, 9-6  etc

`./test_subs  -delay=3 -textdur=3  -mode=udp`

Generate using udp, without incrementing timecodes (always 0-3)
`./test_subs  -delay=3 -textdur=3  -inc=false -mode=udp`

Generate using stdout each 3 seconds, timecodes 0-3, without incrementing timecodes, putting output to stdout
`./test_subs  -delay=3 -textdur=3  -inc=false -mode=stdout`

## Makefile commands
`make clean` -   to clean generates subtitles.
`make run` for compile and run.
`make build` for compile source to binary file
