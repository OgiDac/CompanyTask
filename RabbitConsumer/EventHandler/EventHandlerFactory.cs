using RabbitConsumer.EventHandler.Events;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace RabbitConsumer.EventHandler
{
    public class EventHandlerFactory
    {
        public IUserEventHandler CreateEventHandler(UserEventEnvelope userEventEnvelope)
        {
            var options = new JsonSerializerOptions
            {
                PropertyNameCaseInsensitive = true
            };
            return userEventEnvelope.Type switch
            {
                "UserCreated" => userEventEnvelope.Data.Deserialize<UserCreatedEvent>(options),
                "UserUpdated" => userEventEnvelope.Data.Deserialize<UserUpdatedEvent>(options),
                "UserDeleted" => userEventEnvelope.Data.Deserialize<UserDeletedEvent>(options),
                _ => throw new InvalidOperationException($"Unknown event type: {userEventEnvelope.Type}")
            };
        }
    }
}
