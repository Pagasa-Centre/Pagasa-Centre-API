-- +goose Up
-- +goose StatementBegin

-- Production Ministry
UPDATE ministries
SET
    short_description = 'The Production ministry is responsible for transporting, assembling, and storing the church''s assets and equipment.',
    long_description = 'Our church has a dedicated Production Team that is responsible for creating an atmosphere of worship and celebration through decorations and visual effects. The team is comprised of talented volunteers and designers who are passionate about using their gifts to enhance the worship experience. They also help provide decorations and visual effects for special events throughout the year. They are an integral part of our church’s ministry, helping to create a sense of joy and celebration in everything that we do.

If you are interested in joining our Production Team or would like more information, please don’t hesitate to contact your cell leader or click "Apply" below. We would love to have you join our team and help us create beautiful and inspiring worship experiences for our congregation.',
    thumbnail_url = 'https://cdn.prod.website-files.com/6469d76a492ea69c34c883f8/646a99aaf6f8fe6b1f8d3a70_941A2512.JPG'
WHERE name = 'Production Ministry';

-- Children's Ministry
UPDATE ministries
SET
    short_description = 'Our Children''s Ministry is dedicated to nurturing their spiritual growth and helping them discover the love of Jesus Christ.',
    long_description = 'At Pag-Asa Centre, we believe that children are a precious gift from God, and our Children''s Ministry is dedicated to nurturing their spiritual growth and helping them discover the love of Jesus Christ. We are inspired by Jesus'' own words when He emphasized the importance of children in the Kingdom of God, declaring, "Let the little children come to me, and do not hinder them, for the kingdom of heaven belongs to such as these" (Matthew 19:14, NIV).

Our Children''s Ministry, led by a team of dedicated and passionate volunteers, is committed to creating a safe, loving, and engaging environment where children can learn, grow, and develop a strong foundation of faith. Our volunteers are enthusiastic about guiding your little ones in their journey to know, love, and follow Jesus.

If you share our passion for nurturing the spiritual growth of children and helping them build a lasting relationship with God, we invite you to become a part of our Children''s Ministry team. Your unique talents and gifts can make a meaningful difference in the lives of our young ones.

We firmly believe that by investing in our children today, we are shaping the future of our church and community. Together, as a faith community, we can play a vital role in raising the next godly generation. Our mission is to equip children with the knowledge, values, and faith that will guide them throughout their lives.

Whether you''re a parent looking to involve your child in our ministry or an individual with a heart to serve the Lord through the Children''s Ministry, we welcome you with open arms. Come and join us as we embark on this inspiring journey of faith, hope, and love with the youngest members of our church family.

Together, we can make a lasting impact in the lives of our children and help them shine as lights in the world, sharing God''s love and grace with everyone they meet.

Join our Children''s Ministry by clicking the "Apply" button and filling in the form.',
    thumbnail_url = 'https://cdn.prod.website-files.com/6469d76a492ea69c34c883f8/652963f062e57157ec634e80_WhatsApp%20Image%202023-10-05%20at%2015.31.41.jpeg'
WHERE name = 'Children''s Ministry';

-- Media Ministry
UPDATE ministries
SET
    short_description = 'The Media Ministry is the Church''s evangelistic extension that focuses on using media to spread the word of God.',
    long_description = 'The Media Ministry is the Church''s evangelistic extension that focuses on using media to spread the word of God. The technical assistance needed for our worship services, archived teachings, sermon messages, and other events will be taught. You will have the chance to use various forms of technology, media outlets, and social media platforms to effectively share the Good News of Jesus Christ with people all over the world if you choose to serve in the team!

You have lots of options for serving in the media ministry. The Pag-Asa Centre Media team is here to assist you in learning what you''ll need to know while having fun!

Please contact Pastor Gian or your cell leader if you want to be a member. You can also click "Apply" below and fill in the form!',
    thumbnail_url = 'https://cdn.prod.website-files.com/6469d76a492ea69c34c883f8/646a956d04299d74ab6bb9cb_292101691_5213119612109176_1999081090223236391_n.jpg'
WHERE name = 'Media Ministry';

-- Creative Arts Ministry
UPDATE ministries
SET
    short_description = 'Our Creative Arts Ministry is dedicated to creating a vibrant space where these gifts can flourish, and where we can collectively use them to glorify the Lord.',
    long_description = 'At Pag-Asa Centre, we believe that every individual is uniquely gifted by God with an array of talents and abilities. Our Creative Arts Ministry is dedicated to creating a vibrant space where these gifts can flourish, and where we can collectively use them to glorify the Lord.

Discover Your Creative Calling at Pagasa Centre: Creative Arts Ministry
"God has given each of you a gift from his great variety of spiritual gifts. Use them well to serve one another." – 1 Peter 4:10 (NLT)

The Power of Creative Expression
Our Creative Arts Ministry is all about the beautiful fusion of faith and artistic expression. We find inspiration in the words of 1 Peter 4:10, which remind us that our spiritual gifts are meant to be shared to serve one another and bring glory to God. Through the captivating mediums of dance and acting, we aim to spread joy, love, and our heartfelt praise through the powerful language of movement.

Meet Our Visionary Leader
Under the expert guidance of our visionary leader, Nathan Gordon, a seasoned professional dancer and choreographer, our ministry gathers weekly to craft and refine church-wide productions that inspire, uplift, and engage. Nathan''s passion for the arts and unwavering faith in God''s grace infuse our ministry with boundless creativity and enthusiasm.

Open Doors, Open Hearts
Whether you have been blessed with a creative gift that''s ready to shine or if you''re simply intrigued and eager to explore your artistic side, the doors of our Creative Arts Ministry are wide open to all. We believe that creativity knows no bounds, and everyone is welcome to join us on this journey of discovery, expression, and faith.

Join Us Today
Come and be a part of a community where your creativity is celebrated, and your artistic journey is nurtured. At Pag-Asa Centre''s Creative Arts Ministry, we are more than a team; we are a family united by our love for God and our passion for creativity. Together, we seek to create meaningful, impactful, and spiritually uplifting productions that touch hearts and inspire souls.

Discover, Create, Worship
Unlock the door to your creative potential, and let your talents shine as a vibrant part of our Creative Arts Ministry. Together, we will illuminate the world with the message of God''s love and grace through the universal language of art and expression.

Join us and let your creative spirit soar at Pagasa Centre''s Creative Arts Ministry by clicking "Apply" and filling in the form.',
    thumbnail_url = 'https://cdn.prod.website-files.com/6469d76a492ea69c34c883f8/652964726174b25219845220_WhatsApp%20Image%202023-10-05%20at%2015.31.39.jpeg'
WHERE name = 'Creative Arts Ministry';

-- Music Ministry
UPDATE ministries
SET
    short_description = 'Through music, we have the privilege to lead ourselves and others into a profound connection with the Divine.',
    long_description = '“In the music ministry, we are called to offer our bodies as a living sacrifice, holy and pleasing to God (Romans 12:1). This act of consecration resonates deeply with the essence of worship described in Scripture: ''When the true worshipers will worship the Father in spirit and truth, for the Father is seeking such people to worship Him. God is spirit, and those who worship Him must worship in spirit and truth.'' (John 4:22-24)

As musicians and worshipers, our role is to create an environment where the and truth of God are palpable. Through music, we have the privilege to lead ourselves and others into a profound connection with the Divine. Our songs and melodies become a conduit through which we offer ourselves as living sacrifices wholly dedicated to the Lord''s purpose.

In this ministry, we understand that genuine worship goes beyond external performances; it delves into the depths of our spirits and aligns our hearts with the truth of God''s Word. We are not just singers and instrumentalists; we are worshipers who seek to honour God by worshiping Him in spirit and truth, offering our very selves as pleasing offerings to the Father. As we engage in this true and proper worship, we answer the divine call to seek and serve Him with all our hearts, all our souls, and all our musical talents."',
    thumbnail_url = 'https://cdn.prod.website-files.com/6469d76a492ea69c34c883f8/646a97ba5264c7470ef402f5_293133648_5213119468775857_7202805461124800000_n.jpg'
WHERE name = 'Music Ministry';

-- Ushering & Security Ministry
UPDATE ministries
SET
    short_description = 'The Ushers are the first representative of Jesus Christ for a worship service.',
    long_description = 'As Jesus'' first representatives at worship, we embody spiritual virtues - Holy Spirit''s gifts, service, generosity, and grace. Mentally, physically, emotionally, and spiritually prepared, we attentively oversee services, assisting with various needs. Our duties include creating a warm atmosphere, guiding newcomers to their seats, welcoming and tracking VIP visitors, managing offerings and prayer requests, and expressing gratitude to attendees post-service.

Please contact Sister Bhing Oliveros or your cell leader if you want to be a member. You can also fill in the form by click "Apply" below.',
    thumbnail_url = 'https://cdn.prod.website-files.com/6469d76a492ea69c34c883f8/65296864b3b16041a683d071_WhatsApp%20Image%202023-10-05%20at%2015.31.45.jpeg'
WHERE name = 'Ushering & Security Ministry';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Note: Cannot revert data updates easily without backup, so we’ll just NULL them
UPDATE ministries
SET
    short_description = NULL,
    long_description = NULL,
    thumbnail_url = NULL
WHERE name IN (
               'Production Ministry',
               'Children''s Ministry',
               'Media Ministry',
               'Creative Arts Ministry',
               'Music Ministry',
               'Ushering & Security Ministry'
    );

-- +goose StatementEnd