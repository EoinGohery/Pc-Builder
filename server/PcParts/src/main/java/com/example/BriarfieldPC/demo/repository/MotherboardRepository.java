package com.example.BriarfieldPC.demo.repository;

import com.example.BriarfieldPC.demo.entity.Motherboard;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface MotherboardRepository extends JpaRepository<Motherboard,Long> {
    List<Motherboard> findAll();
}
