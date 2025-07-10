using RabbitConsumer.EventHandler;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System.Text;
using System.Text.Json;

class Program
{
    static async Task Main(string[] args)
    {
        Console.WriteLine("Starting RabbitMQ Consumer...");
        var rabbitUrl = Environment.GetEnvironmentVariable("RABBITMQ_URL")
                 ?? "amqp://guest:guest@localhost:5672/";

        var factory = new ConnectionFactory
        {
            Uri = new Uri(rabbitUrl)
        };


        using var connection = await factory.CreateConnectionAsync();
        using var channel = await connection.CreateChannelAsync();

        string queueName = "user-queue";

        await channel.QueueDeclareAsync(
            queue: queueName,
            durable: false,
            exclusive: false,
            autoDelete: false,
            arguments: null
        );

        Console.WriteLine($"[*] Waiting for messages in '{queueName}'. To exit press CTRL+C");

        var consumer = new AsyncEventingBasicConsumer(channel);
        var eventHandlerFactory = new EventHandlerFactory();
        var options = new JsonSerializerOptions
        {
            PropertyNameCaseInsensitive = true
        };
        consumer.ReceivedAsync += async (model, ea) =>
        {
            var body = ea.Body.ToArray();
            var message = Encoding.UTF8.GetString(body);

            try
            {
                var envelope = JsonSerializer.Deserialize<UserEventEnvelope>(message, options);
                if (envelope != null)
                {
                    var handler = eventHandlerFactory.CreateEventHandler(envelope);
                    var messageToDisplay = handler.HandleEvent();
                    Console.WriteLine(messageToDisplay);
                }
                else
                {
                    Console.WriteLine("[!] Failed to parse envelope.");
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine($"[!] Error handling message: {ex.Message}");
            }

            await Task.Yield();
        };

        await channel.BasicConsumeAsync(
            queue: queueName,
            autoAck: true,
            consumer: consumer
        );

        Console.WriteLine("Press CTRL + C to exit.");
        await Task.Delay(Timeout.Infinite);
    }
}
