1. Assinatura da Função

```go
func <name>() <return> {}
func NewTaskList() *TaskList {}
```

2. Ponteiro

```go
*TaskList   // Ponteiro para a variável
TaskList    // Variável
```

3. Struct Literal

```go
&TaskList{} // Ampersand para criar um ponteiro para o struct
```

4. Inicialização de Slice

```go
make([]<type>, <length>, <capacity>) // Cria um slice vazio com o tipo, tamanho e capacidade especificados
make([]Task, 0, 10) // Cria um slice vazio com o tipo Task, tamanho 0 e capacidade 10
```

```go
[]<type>{} // Cria um slice vazio com o tipo especificado
[]Task{} // Cria um slice vazio com o tipo Task
```

5. Contructor Pattern

```go
func New<type>() <return> {
    return &<type>{}
}

func NewTaskList() *TaskList {
    return &TaskList{
        Tasks: make([]Task, 0),
        NextID: 1,
    }
}
```

6. Inicialização Zero vs Explícita

```go
// Zero - automático
var <variable> <type>
var tl TaskList
// Tasks = nil, NextID = 0
```

```go
// Explícita
var <variable> = <type>{}
var tl = TaskList{}
// Tasks = [], NextID = 1
```

7. Iteração sobre Slice/Array

```go
for i := range <slice> {}
for i := range <array> {}
```

8. Remoção da Array/Slice

```go
<slice> = append(<slice>[:i], <slice>[i+1:]...)
```

```go
<slice>[:i] // Pega todos elementos antes do índice i
<slice>[i+1:] // Pega todos elementos depois do índice i
append(<slice>[:i], <slice>[i+1:]...) // Concatena os dois slices
```

