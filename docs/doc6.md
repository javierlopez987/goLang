# GoLang
## CRUD
## TDD
Test-Driven Development
Técnica de programación que "comienza a trabajar con la funcionalidad sin tenerla hecha" (Test First Development)
Luego, "se van llenando esos huecos", se refactoriza y se va completando el código (Refactoring)

### Update
Mi alternativa previo a ver el código sugerido fue:
    ```js
        // Update ...
        func (p Playlist) Update(s Song)  {
            if p.FindByID(s.ID) != nil {
                p.Add(s)
            }
        }
    ```
El código sugerido fue:
    ```js
        // Update ...
        func (p Playlist) Update(s Song)  {
            p.songs[s.ID] = &s
        }
    ```