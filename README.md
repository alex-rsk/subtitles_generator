Emulate generation of VTT subitles.

Command line parameters

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
