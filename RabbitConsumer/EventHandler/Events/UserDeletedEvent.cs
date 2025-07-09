using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace RabbitConsumer.EventHandler.Events
{
    public record UserDeletedEvent(uint Id) : IUserEventHandler
    {
        public string HandleEvent()
        {
            return $"[Handled] User with the {Id} deleted";
        }
    }

}
