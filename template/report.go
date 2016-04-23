package template

const Report string = `<html style="font-family: monospace;">
   <head>
      <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=0.6, maximum-scale=1.0, user-scalable=no">
      <title>Back up bandit report</title>
   </head>
   <body>
      <style type="text/css">
          html {
              font-family: monospace;
          }

          table td tr h1 h2 h3 h4 h5 {
              margin: 0px;
              padding: 0px;
          }

          .rad {
              border-radius: 3px;
              padding: 3px;
          }

          h1 {
              margin-left: 5px;
          }

          h2 {
              margin-left: 15px;
          }

          h2 {
              margin-left: 25px;
          }

          h3 {
              margin-left: 35px;
          }

          table {
              align-content: left;
              text-align: left;
              border-collapse: collapse;
          }

          .critical {
              background-color: #212121;
              color: white;
          }

          .error {
              background-color: #FF3D00;
              color: white;
          }

          .warning {
              background-color: #FF9100;
              color: white;
          }

          .success {
              background-color: #689F38;
              color: white;
          }

          .tools {
              background-color: #1565C0;
              color: white;
          }
      </style>
      <table width="100%" style="align-content: left;text-align: left;border-collapse: collapse;">
         <tr>
            <td>
               <table id="top-message" width="100%" style="align-content: left;text-align: left;border-collapse: collapse;">
                   <tr><h1 style="margin-left: 5px;">Report {{.Date}}</h1></tr>
               </table>
               <!-- top message -->
               <table width="100%" style="align-content: left;text-align: left;border-collapse: collapse;">
                   <tr>
                       <td>
                        <h2 class="critical rad" style="margin-left: 25px;border-radius: 3px;padding: 3px;background-color: #212121;color: white;">Critical</h2>
                        {{range $critical := .Critical}}
                            <h3 style="margin-left: 35px;">{{$critical}}</h3>
                        {{end}}
                       </td>
                   </tr>
                   <tr>
                       <td>
                        <h2 class="error rad" style="margin-left: 25px;border-radius: 3px;padding: 3px;background-color: #FF3D00;color: white;">Error</h2>
                        {{range $error := .Error}}
                            <h3 style="margin-left: 35px;">{{$error}}</h3>
                        {{end}}
                       </td>
                   </tr>
                   <tr>
                       <td>
                        <h2 class="warning rad" style="margin-left: 25px;border-radius: 3px;padding: 3px;background-color: #FF9100;color: white;">Warning</h2>
                        {{range $warning := .Warning}}
                            <h3 style="margin-left: 35px;">{{$warning}}</h3>
                        {{end}}
                       </td>
                   </tr>
                   <tr>
                       <td>
                        <h2 class="success rad" style="margin-left: 25px;border-radius: 3px;padding: 3px;background-color: #689F38;color: white;">Successful</h2>
                        {{range $success := .Success}}
                            <h3 style="margin-left: 35px;">{{$success}}</h3>
                        {{end}}
                       </td>
                   </tr>
               </table>
               <!-- main -->
               <table width="100%" style="align-content: left;text-align: left;border-collapse: collapse;">
                 <tr>
                   <h1 style="margin-left: 5px;">Tools log</h1>
                 </tr>
                 {{range .Tools}}
                 <h2 class="tools rad" style="margin-left: 25px;border-radius: 3px;padding: 3px;background-color: #1565C0;color: white;">{{.Path}}</h2>
                 <table width="100%" style="align-content: left;text-align: left;border-collapse: collapse;">
                     <tr>
                         {{range .Groups}}
                            <th><h3 style="margin-left: 35px;">{{.}}</h3></th>
                         {{end}}
                     </tr>
                     {{range .ParseOut}}
                        <tr>
                            {{range .}}
                            <td>
                                <h3 style="margin-left: 35px;">{{.}}</h3>
                            </td>
                            {{end}}
                        </tr>
                     {{end}}
                 </table>
                 {{end}}
               </table>
               <!-- bottom message -->
         </td></tr>

      </table>
      <!-- wrapper -->
   </body>
</html>`
