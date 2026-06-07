import Editor, { type OnMount } from '@monaco-editor/react'
import { useTheme } from '@/context/ThemeContext'
import { getLanguage } from '@/lib/codingLanguages'

interface CodeEditorProps {
  language: string
  value: string
  onChange: (value: string) => void
  height?: string
}

export function CodeEditor({
  language,
  value,
  onChange,
  height = '420px',
}: CodeEditorProps) {
  const { resolved } = useTheme()
  const lang = getLanguage(language)

  const handleMount: OnMount = (editor) => {
    editor.focus()
  }

  return (
    <div className="overflow-hidden rounded-xl border border-ink-200/80 dark:border-ink-700">
      <Editor
        height={height}
        language={lang.monacoId}
        value={value}
        onChange={(v) => onChange(v ?? '')}
        onMount={handleMount}
        theme={resolved === 'dark' ? 'vs-dark' : 'light'}
        options={{
          minimap: { enabled: false },
          fontSize: 14,
          lineNumbers: 'on',
          scrollBeyondLastLine: false,
          automaticLayout: true,
          tabSize: 4,
          wordWrap: 'on',
          padding: { top: 12 },
        }}
      />
    </div>
  )
}
