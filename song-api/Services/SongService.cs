using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.Extensions.Caching.Memory;
using Songs.API.Domain.Models;
using Songs.API.Domain.Repositories;
using Songs.API.Domain.Services;
using Songs.API.Domain.Services.Communication;
using Songs.API.Infrastructure;

namespace Songs.API.Services
{
    public class SongService : ISongService
    {
        private readonly ISongRepository _songRepository;
        private readonly IMemoryCache _cache;

        public SongService(ISongRepository songRepository, IMemoryCache cache)
        {
            _songRepository = songRepository;
            _cache = cache;
        }

        public async Task<IEnumerable<Song>> ListAsync()
        {
            // Here I try to get the songs list from the memory cache. If there is no data in cache, the anonymous method will be
            // called, setting the cache to expire one minute ahead and returning the Task that lists the songs from the repository.
            var songs = await _cache.GetOrCreateAsync(CacheKeys.SongsList, (entry) => {
                entry.AbsoluteExpirationRelativeToNow = TimeSpan.FromMinutes(1);
                return _songRepository.ListAsync();
            });
            
            return songs;
        }

    }
}
