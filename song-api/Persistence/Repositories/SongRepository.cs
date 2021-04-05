using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using Songs.API.Domain.Models;
using Songs.API.Domain.Repositories;
using Songs.API.Persistence.Contexts;

namespace Songs.API.Persistence.Repositories
{
    public class SongRepository : BaseRepository, ISongRepository
    {
        public SongRepository(AppDbContext context) : base(context) { }

        public async Task<IEnumerable<Song>> ListAsync()
        {
            return await _context.Songs
                                 .AsNoTracking()
                                 .ToListAsync();

            // AsNoTracking tells EF Core it doesn't need to track changes on listed entities. Disabling entity
            // tracking makes the code a little faster
        }

    }
}