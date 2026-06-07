export interface CodingLanguage {
  id: string
  label: string
  monacoId: string
  supported: boolean
  defaultCode: string
}

export const CODING_LANGUAGES: CodingLanguage[] = [
  {
    id: 'python',
    label: 'Python',
    monacoId: 'python',
    supported: true,
    defaultCode: '# Write your solution here\n',
  },
  {
    id: 'javascript',
    label: 'JavaScript',
    monacoId: 'javascript',
    supported: false,
    defaultCode: '// Coming soon\n',
  },
  {
    id: 'java',
    label: 'Java',
    monacoId: 'java',
    supported: false,
    defaultCode: '// Coming soon\n',
  },
  {
    id: 'cpp',
    label: 'C++',
    monacoId: 'cpp',
    supported: false,
    defaultCode: '// Coming soon\n',
  },
]

export function getLanguage(id: string) {
  return CODING_LANGUAGES.find((l) => l.id === id) ?? CODING_LANGUAGES[0]
}
