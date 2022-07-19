const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');

// Connection URL
const url = 'mongodb://localhost:27017';

// Database Name
const dbName = 'sapi';

// Create a new MongoClient
const client = new MongoClient(url);

mongo();

async function mongo() {
    await client.connect();
    console.log('Connected successfully to server');
    const db = client.db(dbName);

    let r = null;
    r = await db
        .collection('resources')
        .find(q1)
        .toArray();
    console.log(r);

    r = await db
        .collection('resources')
        .find(q2)
        .toArray();
    console.log(r);

    r = await db
        .collection('resources')
        .find(q3)
        .toArray();
    console.log(r);

    console.log('Closing connection');
    client.close();
}

const q1 = {
    tenantId: 'tenant_3',
    'permissions.users': {$in: ['user_7']}
};

const q2 = {
    tenantId: 'tenant_2',
    $or: [
        {'permissions.users': {$in: ['user_9']}},
        {'permissions.groups': {$in: ['group_6']}}
    ]
};

const q3 = {
    tenantId: 'tenant_3',
    $or: [{'permissions.users': {$in: ['user_5']}}, {'permissions.groups': {$in: []}}]
};

// pg.resources.find({
//   tenantId: 'tenant_3',
//   'permissions.account': { $in: ['user_7'] }
// });

// pg.resources.find({
//   tenantId: 'tenant_2',
//   $or: [
//     { 'permissions.account': { $in: ['user_9'] } },
//     { 'permissions.groups': { $in: ['group_6'] } }
//   ]
// });
