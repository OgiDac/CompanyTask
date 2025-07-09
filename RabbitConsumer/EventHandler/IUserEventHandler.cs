using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace RabbitConsumer.EventHandler
{
    public interface IUserEventHandler
    {
        string HandleEvent();
    }
}
