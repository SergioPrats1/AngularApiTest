using AutoMapper;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Songs.API.Controllers.Config;
using Songs.API.Domain.Repositories;
using Songs.API.Domain.Services;
using Songs.API.Extensions;
using Songs.API.Persistence.Contexts;
using Songs.API.Persistence.Repositories;
using Songs.API.Services;

namespace Songs.API
{
    public class Startup
    {
        public IConfiguration Configuration { get; }

        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public void ConfigureServices(IServiceCollection services)
        {
            services.AddMemoryCache();

            services.AddCustomSwagger();

            services.AddControllers().ConfigureApiBehaviorOptions(options =>
            {
                // Adds a custom error response factory when ModelState is invalid
                options.InvalidModelStateResponseFactory = InvalidModelStateResponseFactory.ProduceErrorResponse;
            });

            services
                .AddCors(options =>
                {
                    options.AddPolicy("AllowAll",
                        builder => builder
                        .AllowAnyOrigin()
                        .AllowAnyMethod()
                        .AllowAnyHeader()
                        );

                    options.AddPolicy("signalr",
                        builder => builder
                        .AllowAnyMethod()
                        .AllowAnyHeader()

                        .AllowCredentials()
                        .SetIsOriginAllowed(hostName => true));
                });

            services.AddDbContext<AppDbContext>(options =>
            {
				options.UseSqlite(Configuration.GetConnectionString("sqlite"));
            });

            services.AddScoped<ISongRepository, SongRepository>();

            services.AddScoped<ISongService, SongService>();

            services.AddAutoMapper(typeof(Startup));

            services.AddMvc();
        }

        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseCustomSwagger();

            app.UseRouting();

            app.UseCors("AllowAll");

            app.UseAuthorization();

            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
            });
        }
    }
}