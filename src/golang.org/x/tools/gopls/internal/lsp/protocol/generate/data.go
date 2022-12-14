// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.19
// +build go1.19

package main

// various data tables

// methodNames is a map from the method to the name of the function that handles it
var methodNames = map[string]string{
	"$/cancelRequest":                        "CancelRequest",
	"$/logTrace":                             "LogTrace",
	"$/progress":                             "Progress",
	"$/setTrace":                             "SetTrace",
	"callHierarchy/incomingCalls":            "IncomingCalls",
	"callHierarchy/outgoingCalls":            "OutgoingCalls",
	"client/registerCapability":              "RegisterCapability",
	"client/unregisterCapability":            "UnregisterCapability",
	"codeAction/resolve":                     "ResolveCodeAction",
	"codeLens/resolve":                       "ResolveCodeLens",
	"completionItem/resolve":                 "ResolveCompletionItem",
	"documentLink/resolve":                   "ResolveDocumentLink",
	"exit":                                   "Exit",
	"initialize":                             "Initialize",
	"initialized":                            "Initialized",
	"inlayHint/resolve":                      "Resolve",
	"notebookDocument/didChange":             "DidChangeNotebookDocument",
	"notebookDocument/didClose":              "DidCloseNotebookDocument",
	"notebookDocument/didOpen":               "DidOpenNotebookDocument",
	"notebookDocument/didSave":               "DidSaveNotebookDocument",
	"shutdown":                               "Shutdown",
	"telemetry/event":                        "Event",
	"textDocument/codeAction":                "CodeAction",
	"textDocument/codeLens":                  "CodeLens",
	"textDocument/colorPresentation":         "ColorPresentation",
	"textDocument/completion":                "Completion",
	"textDocument/declaration":               "Declaration",
	"textDocument/definition":                "Definition",
	"textDocument/diagnostic":                "Diagnostic",
	"textDocument/didChange":                 "DidChange",
	"textDocument/didClose":                  "DidClose",
	"textDocument/didOpen":                   "DidOpen",
	"textDocument/didSave":                   "DidSave",
	"textDocument/documentColor":             "DocumentColor",
	"textDocument/documentHighlight":         "DocumentHighlight",
	"textDocument/documentLink":              "DocumentLink",
	"textDocument/documentSymbol":            "DocumentSymbol",
	"textDocument/foldingRange":              "FoldingRange",
	"textDocument/formatting":                "Formatting",
	"textDocument/hover":                     "Hover",
	"textDocument/implementation":            "Implementation",
	"textDocument/inlayHint":                 "InlayHint",
	"textDocument/inlineValue":               "InlineValue",
	"textDocument/linkedEditingRange":        "LinkedEditingRange",
	"textDocument/moniker":                   "Moniker",
	"textDocument/onTypeFormatting":          "OnTypeFormatting",
	"textDocument/prepareCallHierarchy":      "PrepareCallHierarchy",
	"textDocument/prepareRename":             "PrepareRename",
	"textDocument/prepareTypeHierarchy":      "PrepareTypeHierarchy",
	"textDocument/publishDiagnostics":        "PublishDiagnostics",
	"textDocument/rangeFormatting":           "RangeFormatting",
	"textDocument/references":                "References",
	"textDocument/rename":                    "Rename",
	"textDocument/selectionRange":            "SelectionRange",
	"textDocument/semanticTokens/full":       "SemanticTokensFull",
	"textDocument/semanticTokens/full/delta": "SemanticTokensFullDelta",
	"textDocument/semanticTokens/range":      "SemanticTokensRange",
	"textDocument/signatureHelp":             "SignatureHelp",
	"textDocument/typeDefinition":            "TypeDefinition",
	"textDocument/willSave":                  "WillSave",
	"textDocument/willSaveWaitUntil":         "WillSaveWaitUntil",
	"typeHierarchy/subtypes":                 "Subtypes",
	"typeHierarchy/supertypes":               "Supertypes",
	"window/logMessage":                      "LogMessage",
	"window/showDocument":                    "ShowDocument",
	"window/showMessage":                     "ShowMessage",
	"window/showMessageRequest":              "ShowMessageRequest",
	"window/workDoneProgress/cancel":         "WorkDoneProgressCancel",
	"window/workDoneProgress/create":         "WorkDoneProgressCreate",
	"workspace/applyEdit":                    "ApplyEdit",
	"workspace/codeLens/refresh":             "CodeLensRefresh",
	"workspace/configuration":                "Configuration",
	"workspace/diagnostic":                   "DiagnosticWorkspace",
	"workspace/diagnostic/refresh":           "DiagnosticRefresh",
	"workspace/didChangeConfiguration":       "DidChangeConfiguration",
	"workspace/didChangeWatchedFiles":        "DidChangeWatchedFiles",
	"workspace/didChangeWorkspaceFolders":    "DidChangeWorkspaceFolders",
	"workspace/didCreateFiles":               "DidCreateFiles",
	"workspace/didDeleteFiles":               "DidDeleteFiles",
	"workspace/didRenameFiles":               "DidRenameFiles",
	"workspace/executeCommand":               "ExecuteCommand",
	"workspace/inlayHint/refresh":            "InlayHintRefresh",
	"workspace/inlineValue/refresh":          "InlineValueRefresh",
	"workspace/semanticTokens/refresh":       "SemanticTokensRefresh",
	"workspace/symbol":                       "Symbol",
	"workspace/willCreateFiles":              "WillCreateFiles",
	"workspace/willDeleteFiles":              "WillDeleteFiles",
	"workspace/willRenameFiles":              "WillRenameFiles",
	"workspace/workspaceFolders":             "WorkspaceFolders",
	"workspaceSymbol/resolve":                "ResolveWorkspaceSymbol",
}
