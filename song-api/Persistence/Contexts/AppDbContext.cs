using Microsoft.EntityFrameworkCore;
using Songs.API.Domain.Models;

namespace Songs.API.Persistence.Contexts
{
    public class AppDbContext : DbContext
    {
        public DbSet<Song> Songs { get; set; }

        public AppDbContext(DbContextOptions<AppDbContext> options) : base(options) { }

        protected override void OnModelCreating(ModelBuilder builder)
        {
            base.OnModelCreating(builder);
            
            builder.Entity<Song>().ToTable("Song");
            builder.Entity<Song>().HasKey(p => p.Id);
            builder.Entity<Song>().Property(p => p.Id).IsRequired().ValueGeneratedOnAdd();//.HasValueGenerator<InMemoryIntegerValueGenerator<int>>();
            builder.Entity<Song>().Property(p => p.Title).HasMaxLength(80);
            builder.Entity<Song>().Property(p => p.Artist).HasMaxLength(70);

            /*builder.Entity<Song>().HasData
            (
                new Song { Id = 100, Title = "Test Song", Artist = "Dummy", Year = 2021 }, // Id set manually due to in-memory provider
				new Song { Id = 101, Title = "Test Song2", Artist = "Dummy", Year = 2020 }
            );*/
        }
    }
}