package editor

import (
	"github.com/charmbracelet/bubbles/v2/key"
)

type EditorKeyMap struct {
	SendMessage key.Binding
	OpenEditor  key.Binding
	Newline     key.Binding
	Next        key.Binding
	Previous    key.Binding
	AttachmentKeyMap
}

func DefaultEditorKeyMap() EditorKeyMap {
	return EditorKeyMap{
		SendMessage: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "send"),
		),
		OpenEditor: key.NewBinding(
			key.WithKeys("ctrl+v"),
			key.WithHelp("ctrl+v", "open editor"),
		),
		Newline: key.NewBinding(
			key.WithKeys("shift+enter", "ctrl+j"),
			// "ctrl+j" is a common keybinding for newline in many editors. If
			// the terminal supports "shift+enter", we substitute the help text
			// to reflect that.
			key.WithHelp("ctrl+j", "newline"),
		),
		Next: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "down"),
		),
		Previous: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "up"),
		),
		AttachmentKeyMap: AttachmentKeyMaps,
	}
}

// KeyBindings implements layout.KeyMapProvider
func (k EditorKeyMap) KeyBindings() []key.Binding {
	return []key.Binding{
		k.SendMessage,
		k.OpenEditor,
		k.Newline,
		k.Next,
		k.Previous,
		AttachmentKeyMaps.AddFile,
		AttachmentKeyMaps.AddImage,
		AttachmentKeyMaps.DeleteAttachment,
		AttachmentKeyMaps.Escape,
		AttachmentKeyMaps.DeleteAllAttachments,
	}
}

type AttachmentKeyMap struct {
	AddFile              key.Binding
	AddImage             key.Binding
	DeleteAttachment     key.Binding
	AttachmentDeleteMode key.Binding
	Escape               key.Binding
	DeleteAllAttachments key.Binding
}

// TODO: update this to use the new keymap concepts
var AttachmentKeyMaps = AttachmentKeyMap{
	// TODO Bindings()and Help() both set keybinds for chatPage. This causes new
	// key bindings defined in the keys.go file for page/chat to not be
	// respected. Needs to be fixed before we can properly add keybinds here.
	AddFile: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "add file"),
	),
	AddImage: key.NewBinding(
		key.WithKeys("ctrl+f"),
		key.WithHelp("ctrl+f", "add image"),
	),
	DeleteAttachment: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r+{i}", "delete attachment at index i"),
	),
	Escape: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "cancel delete mode"),
	),
	DeleteAllAttachments: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("ctrl+r+r", "delete all attachments"),
	),
}
