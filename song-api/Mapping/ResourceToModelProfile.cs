using AutoMapper;
using Songs.API.Domain.Models;
using Songs.API.Domain.Models.Queries;
using Songs.API.Resources;

namespace Songs.API.Mapping
{
    public class ResourceToModelProfile : Profile
    {
        public ResourceToModelProfile()
        {
            CreateMap<SaveSongResource, Song>();
        }
    }
}