require 'net/http'
require 'json'

def create_agent(body)
    uri = URI('https://metric-api.newrelic.com/metric/v1')
    headers = {
        'Api-Key'=>'API_KEY',
        'Content-Type' =>'application/json'
    }
    http = Net::HTTP.new(uri.host, uri.port)
    http.use_ssl = true
    response = http.post(uri.path, body, headers)

    puts "response #{response.body}, #{response.code}"
rescue => e
    puts "failed #{e}"
end

def get_body(value, timestamp)
    return [{ 
        "metrics":[{ 
           "name":"AdventOfCode", 
           "type":"gauge", 
           "value":value, 
           "timestamp":timestamp, 
           "attributes":{"day":1, "problem":1, "test":"b"} 
           }]
    }].to_json
end

File.open("input.txt", "r") do |file_handle|
    startTime = 1638346450
    file_handle.each_line do |line|
        lineVal = line.gsub("\n",'')
        body = get_body(lineVal.to_i, startTime)
        puts body
        create_agent(body)
        startTime = startTime + 60
    end
end