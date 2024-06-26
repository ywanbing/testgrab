package common

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
)

var htmlTemplate = `<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>%s</title>
    <style>
        body {
            margin: 0 auto;
            font-family: "YaHei Consolas Hybrid", "Noto Sans", "Helvetica Neue", "Segoe UI", Helvetica, Tahoma, Arial, Geneva, Georgia, Palatino, "Times New Roman", "冬青黑体", "Microsoft YaHei", "微软雅黑", "Microsoft YaHei UI", "WenQuanYi Micro Hei", "文泉驿雅黑", Dengxian, "等线体", STXihei, "华文细黑", "Liberation Sans", "Droid Sans", NSimSun, "新宋体", SimSun, "宋体", "Apple Color Emoji", "Segoe UI Emoji";
            color: #222222;
            line-height: 1.5;
            padding: 16px;
            background-color: #ffffff;
            font-size: 16px;
        }

        h1, h2, h3, h4, h5, h6 {
            color: #222222;
            font-weight: bold;
            margin-top: 20px;
            margin-bottom: 10px;
            padding: 0;
        }

        p, blockquote, ul, ol, dl, table {
            margin: 0.4em 0;
        }

        p {
            padding: 0;
        }

        h1 {
            font-size: 2rem;
        }

        h2 {
            font-size: 1.75rem;
        }

        h3 {
            font-size: 1.4rem;
        }

        h4 {
            font-size: 1.2rem;
        }

        h5 {
            font-size: 1rem;
        }

        h6 {
            font-size: 1rem;
        }

        a {
            color: #0099ff;
            margin: 0;
            padding: 0;
            vertical-align: baseline;
            text-decoration: none;
            word-break: break-word;
        }

        a:hover {
            text-decoration: underline;
        }

        a:visited {
            color: purple;
        }

        ul, ol {
            padding: 0;
            padding-left: 24px;
        }

        li {
            line-height: 24px;
        }

        li ul, li ol {
            margin-left: 16px;
        }

        p, ul, ol {
            font-size: 16px;
            line-height: 24px;
        }

        pre {
            display: block;
            overflow-y: hidden;
            overflow-x: auto;
            -moz-tab-size: 4;
            -o-tab-size: 4;
            tab-size: 4;
        }

        code {
            font-family: "YaHei Consolas Hybrid", Consolas, Monaco, "Andale Mono", Monospace, "Courier New";
            color: #8e24aa;
            word-break: break-word;
        }

        pre code {
            display: block;
            padding-left: 0.5em;
            padding-right: 0.5em;
            color: #222222;
            background-color: #e0e0e0;
            line-height: 1.5;
            font-family: "YaHei Consolas Hybrid", Consolas, Monaco, "Andale Mono", Monospace, "Courier New";
            white-space: pre;
            -moz-tab-size: 4;
            -o-tab-size: 4;
            tab-size: 4;
        }

        aside {
            display: block;
            float: right;
            width: 390px;
        }

        blockquote {
            color: #666666;
            border-left: .5em solid #7a7a7a;
            padding: 0 1em;
        }

        blockquote p {
            color: #666666;
        }

        hr {
            display: block;
            text-align: left;
            margin: 1em 0;
            border: none;
            height: 2px;
            background-color: #999999;
        }

        table {
            padding: 0;
            margin: 1rem 0.5rem;
            border-collapse: collapse;
        }

        table tr {
            border-top: 1px solid #cccccc;
            background-color: #ffffff;
            margin: 0;
            padding: 0;
        }

        table tr:hover {
            background-color: #dadada;
        }

        table tr th {
            font-weight: bold;
            border: 1px solid #cccccc;
            margin: 0;
            padding: 6px 13px;
        }

        table tr td {
            border: 1px solid #cccccc;
            margin: 0;
            padding: 6px 13px;
        }

        table tr th :first-child, table tr td :first-child {
            margin-top: 0;
        }

        table tr th :last-child, table tr td :last-child {
            margin-bottom: 0;
        }

        div.vx-mermaid-graph {
            margin: 16px 0px 16px 0px;
            overflow-y: hidden;
        }

        div.vx-flowchartjs-graph {
            padding: 0px 5px 0px 5px;
            margin: 16px 0px 16px 0px;
            width: fit-content;
            overflow: hidden;
        }

        div.vx-wavedrom-graph {
            padding: 0px 5px 0px 5px;
            margin: 16px 0px 16px 0px;
            width: fit-content;
            overflow: hidden;
        }

        div.vx-plantuml-graph {
            padding: 5px 5px 0px 5px;
            margin: 16px 0px 16px 0px;
            width: fit-content;
            overflow: hidden;
        }

        ::selection {
            background-color: #1976d2;
            color: #ffffff;
        }

        ::-webkit-scrollbar {
            background-color: #f5f5f5;
            width: 14px;
            height: 14px;
            border: none;
        }

        ::-webkit-scrollbar-corner {
            background-color: #f5f5f5;
        }

        ::-webkit-scrollbar-button {
            /* This selector affects the styling of both the up & down and left & right buttons of a scrollbar */
            height: 14px;
            width: 14px;
            background-color: #f5f5f5;
        }

        ::-webkit-scrollbar-button:hover {
            background-color: #dadada;
        }

        ::-webkit-scrollbar-button:active {
            background-color: #c0c0c0;
        }

        ::-webkit-scrollbar-track {
            /* This selector affects the styling of the area in the scrollbar between the two buttons */
            background-color: #f5f5f5;
        }

        ::-webkit-scrollbar-thumb {
            /* This selector affects the styling of draggable element of the scollbar */
            border: none;
            background-color: #cdcdcd;
        }

        ::-webkit-scrollbar-thumb:hover {
            background-color: #c0c0c0;
        }

        ::-webkit-scrollbar-thumb:active {
            background-color: #bbbbbb;
        }

        #vx-content span.vx-search-match {
            color: #222222;
            background-color: #4db6ac;
        }

        #vx-content span.vx-current-search-match {
            color: #222222;
            background-color: #66bb6a;
        }

        .vx-alert{
            background-color: #E2EDF3 !important;
            padding: 12px 24px 12px 30px !important;
            border-radius:0 !important;
            border:none !important;
            border-left: 4px solid #498BA7 !important;
            color: inherit !important;
        }

        .vx-alert::before
        {
            background-color: #498BA7;
            border-radius: 100%%;
            color: #fff;
            content: '!';
            font-family: 'Dosis', 'Source Sans Pro', 'Helvetica Neue', Arial, sans-serif;
            font-size: 14px;
            font-weight: bold;
            left: -12px;
            line-height: 20px;
            position: absolute;
            height: 20px;
            width: 20px;
            text-align: center;
            top: 14px;
        }

        .alert-success,.alert-s{
            border-color: #42B983 !important;
            background-color: #D4EDDA !important;
        }

        .alert-success::before,.alert-s::before{
            background-color: #42B983;
        }

        .alert-warning,.alert-w{
            border-color: #ffa502 !important;
            background-color: #fff3cd !important;
        }

        .alert-warning:before,.alert-w::before{
            background-color: #ffa502;
        }

        .alert-info,.alert-i{
            border-color: #70a1ff !important;
            background-color: #CCE5FF !important;
        }

        .alert-info:before,.alert-i::before{
            background-color: #70a1ff;
        }

        .alert-danger,.alert-d{
            border-color: #ff4757 !important;
            background-color: #F8D7DA !important;
        }

        .alert-danger:before,.alert-d::before{
            background-color: #ff4757;
        }

        .alert-light{
            border-color: #C1C1C1 !important;
            background-color: #f8f8f8 !important;
        }

        .alert-light:before{
            background-color: #C1C1C1;
        }

        .alert-dark{
            border-color: #484848 !important;
            background-color: #F0F0F0 !important;
        }

        .alert-dark:before{
            background-color: #484848;
        }

    </style>
</head>
<body>
%s
</body>
</html>`

func GetHtml(title string, content []byte) (html []byte) {
	var buf bytes.Buffer
	_ = goldmark.Convert(content, &buf)
	return []byte(fmt.Sprintf(htmlTemplate, title, string(buf.Bytes())))
}
