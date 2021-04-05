using System.Collections.Generic;
using System.Threading.Tasks;
using Songs.API.Domain.Models;
using Songs.API.Domain.Services.Communication;

namespace Songs.API.Domain.Services
{
    public interface ISongService
    {
         Task<IEnumerable<Song>> ListAsync();
    }
}