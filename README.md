# Cases Materialize

## Descrição do Projeto
    Cases destinados à empresa Materialize. Foi construida uma API para receber e fazer tratativa de dados, de acordo com o que foi solicitado pela empresa. 

### Funcionalidades do sistema
    - Request feito ao endpoint /config/<N> (POST), onde "N" é a quantidade de 'nós' que devem ser configurados.
        - O valor de "N" informado na rota é tratado e feito validações para checar se é um numérico. Caso não seja informado um valor numérico, o sistema retornará um erro com a menssagem específica.
        - Se o valor informado na rota for um valor numérico, este mesmo será salvo nas variáveis de ambiente.
        - Se a requisição for concluída com sucesso, o sistema retornará um JSON com uma mensagem de sucesso e a quantidade de 'nós' configurados.
    
    - Request feito ao endpoint /process (POST)
        Ao analisar o enunciado, levantei duas maneiras de resolver o que estava em questão. Para não deixar vago, desenvolvi e resolvi das duas formas, e, em
        "ambientes" diferentes (separação por branch).

        - Case 1 (branch: master):
            - Nessa requisição é necessário passar um lista de numeros (em qualquer formato: int, float, string).
            - Caso não seja informado nenhum valor na lista, o sistema retornará um erro, informando que a lista está vazia.
            - Caso na lista haja um valor diferente de um numérico (letras e caracteres), o sistema retornará um JSON informando o valor que está incorreto.
            - Caso não ocorra nenhum erro, o sistema irá realizar a soma de todos os 'nós' paralelamente (cada 'nó' contém a soma dos valores informado na lista).
            - Se a requisição for concluída com sucesso, o sistema retornará um JSON com valor total, devolvido pela soma dos 'nós'.

        - Case 2 (branch: staging):
            - Nessa requisição é necessário passar um lista de numeros (em qualquer formato: int, float, string).
            - Caso não seja informado nenhum valor na lista, o sistema retornará um erro, informando que a lista está vazia.
            - Os valores informados na lista serão divididos igualmente para os 'nós' fornecidos.
            - Caso na lista haja um valor diferente de um numérico (letras e caracteres), o sistema retornará um JSON informando o valor que está incorreto e 
            a qual 'nó' esse valor pertence.
            - Caso não ocorra nenhum erro, o sistema irá realizar a soma de todos os 'nós' paralelamente.
            - Se a requisição for concluída com sucesso, o sistema retornará um JSON com valor total, devolvido pela soma dos 'nós'.


#### Executando a aplicação
    1. Primeiro passo:
        - Faça download do golang.
    2. Segundo passo:
        - No terminal, clone o projeto: https://github.com/PedroMarchesi/caseMaterialize.git
    3. Terceiro passo:
        - Configure suas variáveis de ambiente (GOPATH e GOROOT), caso seu sistema operacional seja o windows.
            - Na variável de ambiente GOPATH, aponte-a para a pasta onde está o seu PROJETO. Na variável de ambiente GOROOT, aponte-a para o local onde foi instalado o Golang.
    4. Quarto passo:
        - Rode o comando 'go get' para instalar todas as dependencias e pacotes utilizados na aplicação
    4. Quinto passo:
        - Feito todos os passos acima, rode o projeto com o comando 'go run main.go'.

##### Explicação
    Foram um informadas duas tratativas de erro no enunciado:
        1. Quando um valor diferente de uma lista de números for passado.
        2. Quando um dos nós retorna erro.

    - Case 1 (branch: master):
        Não consegui levantar e nem simular nenhum tipo de erro que poderia acontecer no segundo caso, com os 'nós'. Pois, todas as tratativas possíveis são feitas antes da execução principal da função.

    - Case 2 (branch: staging):
        Nesse case é retornado além do campo que está inválido, também a qual 'nó' ele pertence.
   

###### Conclusão
    Desde já, agradeço imensamente pela confiança e pela oportunidade que me deram. Espero poder contribuir com a empresa e crescermos juntos!
    Qualquer dúvida estou à disposição,
        Att Pedro Marchesi