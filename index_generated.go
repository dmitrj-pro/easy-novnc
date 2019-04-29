// Code generated by index_generate.go; DO NOT EDIT.

package main

import "html/template"

var indexTMPL = template.Must(template.New("").Parse("<!DOCTYPE html>\n<html lang=\"en\">\n\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">\n    <meta name=\"robots\" content=\"noindex\">\n    <title>noVNC</title>\n    <link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/semantic-ui@2.4.2/dist/semantic.min.css\">\n    <!-- easy-novnc (https://github.com/geek1011/easy-novnc) -->\n    <style>\n        body {\n            background: #f4f4f4;\n        }\n\n        * {\n            box-sizing: border-box;\n        }\n\n        .wrapper {\n            display: flex;\n            align-items: center;\n            justify-content: center;\n            height: 100vh;\n        }\n\n        .connect {\n            display: block;\n            flex: 0 0 auto;\n            margin: 24px auto;\n            width: 100%;\n            max-width: 400px;\n            overflow-y: auto;\n            max-height: 90vh;\n            background: #fff;\n            border: 1px solid #d3d3d3;\n            border-radius: 5px;\n            padding: 24px;\n            box-shadow: 0 2px 6px 0 rgba(0, 0, 0, 0.1);\n        }\n\n        @media only screen and (max-width: 520px) {\n            .connect {\n                flex: 1;\n                margin: 0;\n                height: 100%;\n                min-height: 100%;\n                max-height: 100%;\n                width: 100%;\n                min-width: 100%;\n                max-width: 100%;\n                border: none;\n            }\n        }\n    </style>\n</head>\n\n<body>\n    <div class=\"wrapper\">\n        <div class=\"connect\">\n            <h3 class=\"ui dividing header\">noVNC</h3>\n            <form action=\"/vnc.html\" method=\"GET\" class=\"ui form\">\n                {{if .arbitraryHosts}}\n                {{if .arbitraryPorts}}\n                <div class=\"two fields\">\n                    <div class=\"field\">\n                        <label for=\"host\">Host</label>\n                        <input type=\"text\" id=\"host\" placeholder=\"{{.host}}\" autocomplete=\"off\">\n                    </div>\n                    <div class=\"field\">\n                        <label for=\"port\">Port</label>\n                        <input type=\"number\" id=\"port\" min=\"1\" max=\"65535\" placeholder=\"{{.port}}\" autocomplete=\"off\">\n                    </div>\n                </div>\n                {{else}}\n                <div class=\"field\">\n                    <label for=\"host\">Host</label>\n                    <input type=\"text\" id=\"host\" placeholder=\"{{.host}}\">\n                </div>\n                {{end}}\n                {{end}}\n\n                {{if not .noURLPassword}}\n                <div class=\"field\">\n                    <label for=\"password\">Password</label>\n                    <input type=\"password\" name=\"password\" id=\"password\" placeholder=\"Password (can be specified later)\" autofocus>\n                </div>\n                {{end}}\n\n                <input type=\"hidden\" name=\"path\" id=\"path\" value=\"vnc\">\n                <input type=\"hidden\" name=\"autoconnect\" id=\"autoconnect\" value=\"true\">\n                <input type=\"hidden\" name=\"resize\" id=\"resize\" value=\"scale\">\n                <input class=\"ui button\" type=\"submit\" value=\"Connect\">\n\n                {{if not .basicUI}}\n                <h3 class=\"ui dividing header\">Connection Options</h3>\n\n                <div class=\"two fields\">\n                    <div class=\"inline field\">\n                        <div class=\"ui checkbox\">\n                            <input type=\"checkbox\" name=\"reconnect\" id=\"reconnect\" value=\"true\" checked>\n                            <label for=\"reconnect\">Reconnect automatically</label>\n                        </div>\n                    </div>\n                    <div class=\"inline field\">\n                        <div class=\"ui checkbox\">\n                            <input type=\"checkbox\" name=\"show_dot\" id=\"show_dot\" value=\"true\" checked>\n                            <label for=\"show_dot\">Show dot when no cursor</label>\n                        </div>\n                    </div>\n                </div>\n                <div class=\"two fields\">\n                    <div class=\"inline field\">\n                        <div class=\"ui checkbox\">\n                            <input type=\"checkbox\" name=\"bell\" id=\"bell\" value=\"true\">\n                            <label for=\"bell\">Enable bell</label>\n                        </div>\n                    </div>\n                    <div class=\"inline field\">\n                        <div class=\"ui checkbox\">\n                            <input type=\"checkbox\" name=\"view_only\" id=\"view_only\" value=\"true\" {{if .defaultViewOnly}}checked{{end}}>\n                            <label for=\"view_only\">View only</label>\n                        </div>\n                    </div>\n                </div>\n                {{else}}\n                <input type=\"hidden\" name=\"reconnect\" id=\"reconnect\" value=\"true\">\n                <input type=\"hidden\" name=\"show_dot\" id=\"show_dot\" value=\"true\">\n                <input type=\"hidden\" name=\"bell\" id=\"bell\" value=\"false\">\n                <input type=\"hidden\" name=\"view_only\" id=\"view_only\" value=\"{{if .defaultViewOnly}}true{{else}}false{{end}}\">\n                {{end}}\n            </form>\n        </div>\n    </div>\n    <script>\n        var path = document.getElementById(\"path\");\n        var host = document.getElementById(\"host\");\n        var port = document.getElementById(\"port\");\n\n        function updatePath() {\n            var addr = \"vnc\";\n            if (host && host.value.trim() != \"\") {\n                addr = addr + \"/\" + encodeURIComponent(host.value.trim());\n                if (port && port.value.toString().trim() != \"\") {\n                    addr = addr + \"/\" + port.value.toString().trim();\n                }\n            }\n            path.value = addr;\n        }\n\n        if (host) {\n            host.addEventListener(\"input\", updatePath);\n            host.addEventListener(\"keyup\", updatePath);\n            host.addEventListener(\"blur\", updatePath);\n        }\n\n        if (port) {\n            port.addEventListener(\"input\", updatePath);\n            port.addEventListener(\"keyup\", updatePath);\n            port.addEventListener(\"blur\", updatePath);\n        }\n    </script>\n</body>\n\n</html>"))
