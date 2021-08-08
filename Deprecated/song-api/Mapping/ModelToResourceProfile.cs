using AutoMapper;
using Songs.API.Domain.Models;
using Songs.API.Domain.Models.Queries;
using Songs.API.Extensions;
using Songs.API.Resources;

namespace Songs.API.Mapping
{
    public class ModelToResourceProfile : Profile
    {
        public ModelToResourceProfile()
        {
            CreateMap<Song, SongResource>();
        }
    }
}