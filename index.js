require('dotenv').config();

const AWS = require("aws-sdk");
const axios = require('axios').default;
const pino = require('pino');

const logger = pino({ level: process.env.LOG_LEVEL || 'info' });

const zeroDate = "2000-01-01T00:00:00+00:00"

AWS.config.update({
  region: "us-east-1"
});

var baseUrl = "https://api.weather.gov";
var station = "KIND";

async function getObservation() {
  try {
    logger.info('retrieving observation')
    let url = `${baseUrl}/stations/${station}/observations/latest`;
    logger.debug('url: ' + url);
    let record = await axios.get(url, { headers: { 'User-Agent': 'Axios' } });

    if (record.status != 200) {
      logger.error(`not-ok http response (${record.status}): ${record.statusText}`);
      return null;
    }

    logger.debug('got observation: ' + record.data.id);

    return record.data;
  }
  catch (error) {
    logger.error(error);
    return null;
  }
}

async function sendObservation(observation) {
  try {
    logger.info("sending observation");
    let item = AWS.DynamoDB.Converter.marshall(observation);
    item.locationid = { S: observation.properties.station };
    item.observationtime = { S: observation.properties.timestamp };

    var params = {
      TableName: 'weather-data',
      Item: item
    };

    let dynamodb = new AWS.DynamoDB();
    await dynamodb.putItem(params).promise();
  }
  catch (error) {
    logger.error('failed to send observation: ' + error)
  }
}

async function main() {
  let record = await getObservation();
  if (record == null) {
    logger.fatal('failed to get observation');
    return;
  }

  await sendObservation(record);
}

exports.handler = async function(event) {
  return await main();
}