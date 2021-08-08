using System.Collections.Generic;
using System.Threading.Tasks;
using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Cors;
using Songs.API.Domain.Models;
using Songs.API.Domain.Services;
using Songs.API.Resources;

namespace Songs.API.Controllers
{
    [Route("/api/songs")]
    [Produces("application/json")]
    [ApiController]
    public class SongsController : Controller
    {
        private readonly ISongService _songService;
        private readonly IMapper _mapper;

        public SongsController(ISongService songService, IMapper mapper)
        {
            _songService = songService;
            _mapper = mapper;
        }

        /// <summary>
        /// Lists all Songs
        /// </summary>
        /// <returns>List os songs.</returns>
        [EnableCors("AllowAll")]
        [HttpGet]
        [ProducesResponseType(typeof(IEnumerable<SongResource>), 200)]
        public async Task<IEnumerable<SongResource>> ListAsync()
        {
            var songs = await _songService.ListAsync();
            var resources = _mapper.Map<IEnumerable<Song>, IEnumerable<SongResource>>(songs);

            return resources;
        }

    }
}