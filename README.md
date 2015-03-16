# Meechum - A new take at monitoring

Meechum is a three way software,

- each of your servers needs to have a meechum, let's be honest everybody
 needs a guy like that. What will he do ? he'll ask his backend for
 configuration. Like, hi i'm a meetchum from group secretService, and also
 leader the PresidentSEcurityGroup, what should i do ? And he'll do what he asked, no questions. And he'll personnaly report to whatever you want.
 
- A meechumcli, that's a particular meechum to talk to your backend of choise, let's be etcd or Consul, your work, your choice. This cli will give you the ability to override configurations, to edit groups, or checks.
 
- A meechumagregatorbackend, this backend if you chose it (we don't) is aggregating multiples sources of datas (say each of your meechum reports errors at the same time, let's do something about it). That's not clear at the moment. But for example, each of your meechum report to slack for errors, you'll also want them to be store in a postgre backend, he can do that. The aggregator is a second step in what we want to do.

# Meechum

    He can do checks :
        
        {
            "name":     "mysql-check",
            "cmd":      "$PLUGINS/databases/mysql_check.sh -p 3306",
            "every":    "5m",
            "repeat":   "10m"
        }

# Meechumcli

    $ meechumcli --help
        status
            node {id}
            nodes
            check {id}
            checks
            logs
        nodes
            list
            status
        group
            list
            add
            edit
            delete
        check
            list
            add
            edit
            delete

Pretty self explanatory. The cli is just a way to interact against meechum api,
the api is enable is the meechum node is started with --api

# MeechumAPI

    POST /v1/alert

# Meechum - Storage (consul/etcd)

    /v1 - api Version
        /groups
            /cassandra : {"checks":["check-http","check-ntp","check-prout"]}
                
        /checks
            
        /nodes
            /node1/
                status
                specs 
        
    
        /handlers
            /slack: {"config":}