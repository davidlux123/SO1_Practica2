const express = require("express");
const morgan = require("morgan");
const cors = require("cors");
const client = require('./gRPC_client')
let random = require('random-name')

//Settings
const app = express();
app.set('port', process.env.API_PORT || 3000);

//Middls
app.use(express.json());
app.use(morgan("dev"));
app.use(cors());

//endpionts
app.get('/', (req, res) => {
    res.status(200).json({name: random(), msg: 'Hello World :)))'});
})

app.post('/Jugar', (req, res)=>{
    const {game_id, players} = req.body;
    
    client.SendResultGame({game_id, players}, function(err, response) {
        if(err) {
            res.status(400).json('Error on sending result');
            return
        }
        res.status(200).json(response.response_Game)
    });
})

app.listen(app.get('port'), () => {
    console.log('Servidor en el puerto', app.get('port'));
});


