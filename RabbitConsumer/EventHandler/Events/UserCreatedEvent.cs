using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace RabbitConsumer.EventHandler.Events
{
    public record UserCreatedEvent(string Email, string Name) : IUserEventHandler
    {
        public string HandleEvent()
        {
            return $"[Handled] User Created: {Name}, {Email}";
        }
    }

}
