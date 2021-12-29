const express = require('express');
const app = express();
const MongoClient = require('mongodb').MongoClient;
const url = 'mongodb://miel:miguelesdb@34.125.174.190:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false';


const http = require('http');
const socketio =  require('socket.io');
const cors = require('cors');
const servidor = http.createServer(app);
const port = process.env.PORT || 4000;

//----------Configuracion de redis --------
const redis =  require('redis');
const REDIS_PORT = 6379;
const REDIS_HOST = "34.125.174.190" 
const REDIS_PASSWORD = "miguelesdb"
const redisClient = redis.createClient({
    url: `redis://:${REDIS_PASSWORD}@${REDIS_HOST}:${REDIS_PORT}`
});
redisClient.connect();
//----------------------------------------------------

const consultas = {
    'con1' : '',
    'con2' : '',
    'con3' : '',
    'con4' : '',
    'con5' : '',
    'con6' : ''
}

app.use(cors());

const io = socketio(servidor, {
    cors: {
        origin: "*"
    },
});

io.on('connection', socket => {
    interval = setInterval(() => {
        consultaRedis();
        consultaMongo();
        socket.emit("data", consultas);
    }, 2000);
});

const consultaRedis = async () => {
    //CONSULTA 5
    res = await redisClient.lRange('personas', 0, 4);
    let arr = []
    for (txt of res){
        arr.push(JSON.parse(txt))
    }
    consultas['con5'] = arr;
    //CONSULTA 6
    jsonOb = {}
    let rango = ''
    for(i = 1; i < 6; i++){
        rango = "rango" + i
        res = await redisClient.get(rango)
        if(res == null){
            res = 0
        }
        jsonOb[rango] = parseInt(res)
    }
    consultas['con6'] = jsonOb;
    //await redisClient.quit();
};



function consultaMongo(){
    MongoClient.connect(url, function(err, db) {
        if (err) throw err;
        var dbo = db.db("vacunadosData");
        //CONSULTA 1
        dbo.collection("vacunados").find('').toArray(function(err, result) {
            if (err) throw err;
            consultas['con1'] = result
        });
        //CONSULTA 2
        var query = [{$match:{n_dose:2}},{$group:{_id:'$location', total:{$sum:1}}},{ $sort : { total : -1 } }]
        dbo.collection("vacunados").aggregate(query).toArray(function(err, result) {
          if (err) throw err;
          consultas['con2'] = result
        });
        //CONSULTA 3
        var query = [{$match:{n_dose:1}},{$group:{_id:'$location', total:{$sum:1}}}, { $sort : { total : -1 } }]
        dbo.collection("vacunados").aggregate(query).toArray(function(err, result) {
          if (err) throw err;
          consultas['con3'] = result
        });
        //CONSULTA 4
        var query = [
            {$match:{n_dose:2}},
            {$group:{_id:'$location', total:{$sum:1}}},
            { $sort : { total : -1 } }
            ]
        dbo.collection("vacunados").aggregate(query).toArray(function(err, result) {
          if (err) throw err;
          consultas['con4'] = result
          db.close()
        });
    });
}

servidor.listen(4000, () => console.log('Server levantado en el puerto 4000'));
