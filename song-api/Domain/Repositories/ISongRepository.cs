using System.Collections.Generic;
using System.Threading.Tasks;
using Songs.API.Domain.Models;

namespace Songs.API.Domain.Repositories
{
    public interface ISongRepository
    {
        Task<IEnumerable<Song>> ListAsync();
    }
}