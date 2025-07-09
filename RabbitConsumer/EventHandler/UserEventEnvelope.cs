using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace RabbitConsumer.EventHandler
{
    public record UserEventEnvelope(string Type, JsonElement Data);
}
